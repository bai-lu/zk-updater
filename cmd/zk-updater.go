package main

import (
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zk-updater/internal/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := checkenv(); err != nil {
		log.Fatalln("依赖环境检查失败:", err)
	}

	log.Infoln(config.NodeList)
	// config import 隐性执行 init()
	// 开启巡检
	for i := 0; i < len(config.NodeList); i++ {
		node := config.NodeList[i]
		// 每天巡检, 如果节点ZK修改时间超过10天, 通知执行更新
		go node.Inspect(24 * time.Hour)
		// 监听消息队列, 收到通知执行更新
		go node.Update()
	}

	// 主进程
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-c
		log.Infof("收到信号 -- ", s)
		switch s {
		case syscall.SIGHUP:
			log.Infof("收到终端断开信号, 忽略")
		case syscall.SIGINT, syscall.SIGTERM:
			log.Info("应用准备退出")
			return
		}
	}
}

func checkenv() error {
	// 检查JRE
	if os.Getenv("JAVA_HOME") == "" {
		err := errors.New("JRE not found")
		return err
	} else {
		log.Infoln("found JRE:", os.Getenv("JAVA_HOME"))
	}

	// 检查 RECEIVERS
	if os.Getenv("RECEIVERS") == "" {
		err := errors.New("从环境变量读取通知人列表失败")
		return err
	} else {
		log.Infoln("通知人列表:", os.Getenv("RECEIVERS"))
	}
	// 检查集群 & 检查Azure账号挂载文件
	switch os.Getenv("CONTAINER_S_CLUSTER") {
	case "staging":
		_, err := os.Lstat("/etc/conf/xxx/azure-account-staging")
		if err != nil {
			err := errors.New("检查账号挂载文件失败")
			return err
		}
	case "c3":
		_, err := os.Lstat("/etc/conf/xxx/azure-account-azcn")
		if err != nil {
			err := errors.New("检查账号挂载文件失败")
			return err
		}
	case "aws-mb":
		_, err := os.Lstat("/etc/conf/xxx/azure-account-azind")
		if err != nil {
			err := errors.New("检查账号挂载文件失败")
			return err
		}
	default:
		err := errors.New("从环境变量读取集群信息失败")
		return err
	}
	log.Infoln("集群信息:", os.Getenv("CONTAINER_S_CLUSTER"))

	// 检查KeyCenter
	maxTimes := 5
	for i := 0; i < maxTimes; i++ {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", "0000"), time.Second)

		if err != nil {
			log.Errorln("keycenter agent not ready, retry connect...:", err)
		}
		if conn != nil {
			defer conn.Close()
			log.Infoln("keycenter agent is running on ", net.JoinHostPort("127.0.0.1", "0000"))
			break
		}
		if i == maxTimes-1 {
			log.Errorln("keycenter agent not running:", err)
			return err
		}
		time.Sleep(3 * time.Second)

	}

	return nil
}
