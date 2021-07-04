package kcagent

import (
	"net"
	"os"

	"git.n.github.com/golang/thriftlib/com/github/keycenter/agent"
	"github.com/apache/thrift/lib/go/thrift"

	log "github.com/sirupsen/logrus"
)

const (
	HOST = "127.0.0.1"
	PORT = "0000"
)

func Encrypt(sid string, raw []byte, userOnlySecret []byte, compressType agent.CompressionType) (r []byte) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocalFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Errorln(os.Stderr, "encrypt error resolving address:", err)
	}

	useTransport := transportFactory.GetTransport(transport)
	kcClinet := agent.NewKeycenterAgentClientFactory(useTransport, protocalFactory)
	if err = transport.Open(); err != nil {
		log.Errorln(os.Stderr, "Error opening socket to 127.0.0.1:0000", " ", err)
	}
	defer transport.Close()
	r, err1 := kcClinet.Encrypt(sid, raw, userOnlySecret, compressType)
	if err1 != nil {
		log.Errorln(os.Stderr, "encrypt fail:", err1)
	}
	return r
}

func Decrypt(sid string, cipher []byte, userOnlySecret []byte, compressType agent.CompressionType) (r []byte) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocalFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Errorln(os.Stderr, "decrypt error resolving address:", err)
	}

	useTransport := transportFactory.GetTransport(transport)
	kcClinet := agent.NewKeycenterAgentClientFactory(useTransport, protocalFactory)
	if err = transport.Open(); err != nil {
		log.Errorln(os.Stderr, "Error opening socket to 127.0.0.1:0000", " ", err)
	}
	defer transport.Close()
	r, err1 := kcClinet.Decrypt(sid, cipher, userOnlySecret, compressType)
	if err1 != nil {
		log.Errorln(os.Stderr, "decrypt fail:", err1)
	}
	return r
}
