package node

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"zk-updater/internal/notify"
	"zk-updater/internal/utils/kcagent"

	"git.n.github.com/golang/thriftlib/com/github/keycenter/agent"
	"github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
)

var receivers string = os.Getenv("RECEIVERS")

type Node struct {
	Cluster        string        // 集群名称
	Mode           string        // 工作模式
	Server         []string      // zk server
	Path           string        // zk path
	SessionTimeout time.Duration // 会话持续时间
	Context        []byte        // 节点内容
	Stat           *zk.Stat      // 节点状态
	OutDate        chan struct{}
}

// 持续更新SAS状态
func (node *Node) Inspect(interval time.Duration) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln("巡检过程发生异常, 已忽略", err)
		}
	}()
	for {
		nowUnix := int64(time.Now().Unix() * 1000)
		err := node.Reflash()
		if err != nil {
			time.Sleep(interval)
			continue
		}
		log.Infoln("刷新Node状态", "当前时间戳:", nowUnix, "上次更新时间戳:", node.Stat.Mtime)
		// 发送飞书通知

		text := fmt.Sprintf("集群%s的%s token还有%d天的有效期, 发呆中...", node.Cluster, node.Mode, ExpireDay(node.Stat.Mtime))
		msg := notify.Message{
			"receivers": receivers, // , 分隔多个通知人
			"title":     "每日巡逻",
			"text":      text,
			"img":       "img_069e18b5-aa57-4f7c-b8a5-dcf4aaf05bal",
		}
		notify.Push(msg)

		// 如果ZK修改时间超过10天, 通知执行更新
		// 10*24*60*60*1000 = 864000000 毫秒 = 10天
		// 1*60*60*1000 = 3600000 毫秒 = 1 小时
		// 1*60*1000 = 60000 毫秒 = 1 分钟
		if nowUnix-node.Stat.Mtime > 864000000 {
			log.Infoln("检测到token即将过期", "当前时间戳:", nowUnix, "上次更新时间戳:", node.Stat.Mtime)
			node.OutDate <- struct{}{}
		}
		time.Sleep(interval)
	}
}

func (node *Node) Update() {
	for {
		<-node.OutDate
		go node.Upload()
	}
}

func (node *Node) Upload() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln("上传token发生异常, 已忽略", err)
		}
	}()
	title := "更新程序启动"
	var text string
	token, err := node.GenSasToken()
	if err != nil {
		log.Errorln("appassembler生成token失败, 结束更新", node.Cluster, node.Mode)
		text = fmt.Sprintf("集群%s的%s token调用appassembler调用失败了, 更新失败", node.Cluster, node.Mode)
		msg := notify.Message{
			"receivers": receivers, // , 分隔多个通知人
			"title":     title,
			"text":      text,
			"img":       "img_f72bd231-2e4c-4bd3-8d04-2946c0ee9ebl",
		}
		notify.Push(msg)
		return
	}
	// 执行加密&base64
	origin_data := []byte(token)
	var userOnlySecret string = ""
	secret := []byte(userOnlySecret)
	var enum agent.CompressionType = agent.CompressionType_NONE

	encArray := kcagent.Encrypt("xxxx", origin_data, secret, enum)
	secure_str := base64.StdEncoding.EncodeToString([]byte(encArray))
	log.Infoln("生成访问凭证成功", node.Cluster, node.Mode)
	err = node.upload(secure_str)
	if err != nil {
		log.Errorln("上传token失败, 结束更新", node.Cluster, node.Mode)
		text = fmt.Sprintf("集群%s的%s token上传到zookeeper失败了, 更新失败", node.Cluster, node.Mode)
		msg := notify.Message{
			"receivers": receivers, // , 分隔多个通知人
			"title":     title,
			"text":      text,
			"img":       "img_f72bd231-2e4c-4bd3-8d04-2946c0ee9ebl",
		}
		notify.Push(msg)
		return
	}
	log.Infoln("上传访问凭证到ZK成功", node.Cluster, node.Mode)
	// 发送飞书消息

	text = fmt.Sprintf("集群%s的%s token更新成功了, 目前Token还有%d天的有效期", node.Cluster, node.Mode, ExpireDay(node.Stat.Mtime))
	msg := notify.Message{
		"receivers": receivers, // , 分隔多个通知人
		"title":     title,
		"text":      text,
		"img":       "img_81729869-5920-4ed5-a8fd-40137db9e8cl",
	}
	notify.Push(msg)

}

// WorkDir 获取程序运行时根目录
func workDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	wd := filepath.Dir(execPath)
	if filepath.Base(wd) == "bin" {
		wd = filepath.Dir(wd)
	}
	return wd, nil
}

func (node *Node) GenSasToken() (token string, err error) {
	// /bin/bash appassembler/bin/xxxx-sas-generateCtrl.sh --accounts appassembler/bin/" + location + " --mode " + mode
	cluster := node.Cluster
	mode := node.Mode
	appDir, err := workDir()
	if err != nil {
		log.Errorln("获取工作目录失败", err)
	}
	binDir := filepath.Join(appDir, "appassembler/bin/xxxx-sas-generateCtrl.sh")
	var accountFile string
	switch os.Getenv("CONTAINER_S_CLUSTER") {
	case "staging":
		accountFile = "/etc/conf/xxx/azure-account-staging"
	case "c3":
		accountFile = "/etc/conf/xxx/azure-account-azcn"
	case "aws-mb":
		accountFile = "/etc/conf/xxx/azure-account-azind"
	default:
		err = errors.New("读取微软账号文件失败")
		return
	}
	params := []string{
		binDir,
		"--accounts",
		accountFile,
		"--mode",
		mode,
	}
	cmd := exec.Command("/bin/bash", params...)
	cmd.Env = os.Environ()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Errorln("执行appassembler命令失败", cluster, mode, err)
	}
	log.Debugln(cmd.Args)
	if stderr.String() != "" {
		err = errors.New(stderr.String())
		log.Errorln("appassembler生成token过程存在异常输出", cluster, mode, stderr.String())
		return
	}
	token = stdout.String()
	return
}

func (node *Node) Verify() error {
	secure_str := string(node.Context)
	var userOnlySecret string = ""
	secret := []byte(userOnlySecret)
	var enum agent.CompressionType = agent.CompressionType_NONE
	encArray, err := base64.StdEncoding.DecodeString(secure_str)
	if err != nil {
		log.Println("decrypt fail:", err)
		return err
	}
	decArray := kcagent.Decrypt("xxxx", encArray, secret, enum)
	result := string(decArray)
	log.Println("origin:", result)
	return nil
}

func (node *Node) Reflash() error {
	option := zk.WithEventCallback(callback)
	conn, _, err := zk.Connect(node.Server, node.SessionTimeout, option)

	if err != nil {
		log.Errorln("建立SAS ZK连接失败", err)
	}
	node.Context, node.Stat, err = conn.Get(node.Path)
	if err != nil {
		log.Errorln("获取SAS最近修改时间失败", err)
	}
	log.Infof("获取节点状态信息成功 %#v", node.Stat)
	conn.Close()
	return err
}

func (node *Node) upload(key string) error {
	option := zk.WithEventCallback(callback)
	conn, _, err := zk.Connect(node.Server, node.SessionTimeout, option)

	if err != nil {
		log.Errorln("建立SAS ZK连接失败", err)
		return err
	}
	var storageInfo map[string]interface{}
	err = json.Unmarshal(node.Context, &storageInfo)
	if err != nil {
		log.Errorln("解析StorageInfo失败", err)
		return err
	}
	storageInfo["AZURE"] = key
	data, err := json.MarshalIndent(storageInfo, "", "    ")
	if err != nil {
		log.Errorln("序列化StorageInfo失败", err)
		return err
	}
	// node.Context
	node.Stat, err = conn.Set(node.Path, data, node.Stat.Version)
	if err != nil {
		log.Errorln("更新ZK信息失败", err)
		return err
	}
	conn.Close()
	return nil
}

func callback(event zk.Event) {
	log.Infoln("ZK Connect State:", event.State.String())
}

func ExpireDay(mTime int64) int64 {
	nowUnix := int64(time.Now().Unix() * 1000)
	expire := 30*1000*24*60*60 - (nowUnix - mTime)
	day := expire / (1000 * 24 * 60 * 60)
	return day
}
