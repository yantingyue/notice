package api

import (
	"encoding/json"
	"fmt"
	"log"
	"notice/internal/cli"
)

type FeishuReq struct {
	MsgType string `json:"msg_type"`
	Content string `json:"content"`
}
type FeishuResp struct {
	Extra         string `json:"extra"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

var (
	FS = "https://open.feishu.cn/open-apis/bot/v2/hook/5e243fc2-ab6b-4ed7-8960-6a14e391224f"
)

func FeiShuUrl() {
	payload := &FeishuReq{
		MsgType: "text",
		Content: fmt.Sprintf("{\"text\":\"%s\"}", "抢单成功"),
	}
	jsonBytes, _ := json.Marshal(payload)
	_, err := cli.Post(FS, nil, jsonBytes)
	if err != nil {
		log.Println(err)
	}
}
