package kcagent

import (
	"encoding/base64"
	"testing"

	"git.n.github.com/golang/thriftlib/com/github/keycenter/agent"
)

func TestKCAgent(t *testing.T) {
	token := "hello world"
	origin_data := []byte(token)
	var userOnlySecret string = ""
	secret := []byte(userOnlySecret)
	var enum agent.CompressionType = agent.CompressionType_NONE

	encArray := Encrypt("xxxx", origin_data, secret, enum)
	secure_str := base64.StdEncoding.EncodeToString([]byte(encArray))
	t.Log("加密成功", token, secure_str)

}

// INFO[0000] 加密成功 hello world GBCR3SkiZrfzo7m93vv465blGBIdCLae1nZDlaWfWXrtzGgQtAEYEOqvL0w6CB8S0ElIylgwiwQYFHAIg/apfXyNkJFpk07Ox4dD3nK3AA==
// 密文是随变的
