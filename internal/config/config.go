package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"zk-updater/internal/node"

	"github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
)

var NodeList []node.Node

type Config []map[string]string

func init() {
	//解码
	//以只读方式打开config.json
	cluster := os.Getenv("CONTAINER_S_CLUSTER")
	clusterConfig := fmt.Sprintf("%s.json", cluster)
	appDir, err := workDir()
	if err != nil {
		log.Fatalln("获取工作目录失败", err)
	}
	clusterConfigFile := filepath.Join(appDir, clusterConfig)
	file, err := os.Open(clusterConfigFile)
	if err != nil {
		log.Fatalln("打开配置文件失败", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	for decoder.More() {
		err := decoder.Decode(&config)
		if err != nil {
			log.Fatalln("加载配置文件失败", err)
		}
		log.Infoln(config)
	}

	for _, nodeInfo := range config {
		node := node.Node{
			Cluster:        nodeInfo["cluster"],
			Mode:           nodeInfo["mode"],
			Server:         []string{nodeInfo["server"]},
			Path:           nodeInfo["path"],
			SessionTimeout: 5 * time.Second,
			Context:        []byte{},
			Stat:           &zk.Stat{},
			OutDate:        make(chan struct{}),
		}
		NodeList = append(NodeList, node)
	}
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
