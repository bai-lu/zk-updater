package node

import (
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// staging work [tjwqstaging.zk.hadoop.srv:2181] /xxxx/clusters/20005/Configuration/SecureInfo 5s [] 0x14000192000 0x140001100c0}]

var Testnode Node

func initNode() {
	Testnode = Node{
		Cluster:        "staging",
		Mode:           "work",
		Server:         []string{"tjwqstaging.zk.hadoop.srv:2181"},
		Path:           "/xxxx/clusters/20005/Configuration/SecureInfo",
		SessionTimeout: 5 * time.Second,
		Context:        []byte{},
		Stat:           &zk.Stat{},
		OutDate:        make(chan struct{}),
	}

}

func TestNode(t *testing.T) {
	initNode()
	token, err := Testnode.GenSasToken()
	if err != nil {
		t.Log(err)
	}
	t.Log(token)

}

func TestExpire(t *testing.T) {
	if int64(time.Now().Unix()*1000)-Testnode.Stat.Mtime > 10*24*60*60*1000 {
		t.Log("即将过期")
	}
}
