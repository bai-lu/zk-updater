package notify

// 发送消息到lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"

	"zk-updater/internal/utils/httpclient"

	log "github.com/sirupsen/logrus"
)

type Lark struct{}

type LarkAuthInfo struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type LarkAuthResp struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccesstoken string `json:"tenant_access_token"`
}

func (lark *Lark) authLark() string {
	authUrl := "https://open.f.***.cn/open-apis/auth/v3/tenant_access_token/internal/"

	larkAuthInfo := LarkAuthInfo{
		AppId:     "",
		AppSecret: "",
	}
	body, err := json.Marshal(larkAuthInfo)
	if err != nil {
		log.Errorf("序列化飞书认证信息失败", err)
	}
	responseWrapper := httpclient.PostJson(authUrl, string(body), 3)
	larkAuthResp := new(LarkAuthResp)
	json.Unmarshal([]byte(responseWrapper.Body), larkAuthResp)
	return larkAuthResp.TenantAccesstoken
}

func (lark *Lark) Send(msg Message) {
	tenantAccesstoken := lark.authLark()
	larkUrl := "https://open.f.***.cn/open-apis/message/v4/send/"
	receivers := strings.Split(msg["receivers"], ",")
	for _, receiver := range receivers {
		data := parseTemplate(msg, receiver)
		payload := strings.NewReader(data)
		lark.send(payload, larkUrl, tenantAccesstoken)
	}
}

func (lark *Lark) send(payload *strings.Reader, larkUrl string, tenantAccesstoken string) {
	timeout := 30
	maxTimes := 3
	i := 0
	for i < maxTimes {
		resp := httpclient.PostLark(larkUrl, payload, timeout, tenantAccesstoken)
		if resp.StatusCode == 200 {
			log.Infoln("lark#发送飞书消息成功#%s#消息内容-%s", resp.Body)
			break
		}
		i += 1
		time.Sleep(3 * time.Second)
		if i < maxTimes {
			log.Errorf("lark#发送消息失败#%s#消息内容-%s", resp.Body)
		}
	}
}

func parseTemplate(msg Message, receiver string) string {
	tmpl, err := template.New("notify").Parse(Template)
	if err != nil {
		return fmt.Sprintf("解析通知模板失败: %s", err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]interface{}{
		"receiver": receiver,
		"title":    msg["title"],
		"text":     msg["text"],
		"img":      msg["img"],
	})
	return buf.String()
}
