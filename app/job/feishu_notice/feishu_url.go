package feishu_notice

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
	tokenUrlMap = map[string][]string{
		Token1: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/06a68041-bcdf-4f50-86d0-dd5a9a6012ed",
			"https://open.feishu.cn/open-apis/bot/v2/hook/186251cb-e289-4f5e-a06c-a708706c8eb9",
			"https://open.feishu.cn/open-apis/bot/v2/hook/1d35df52-2e54-45db-89f7-3a7383dd3958",
		}, //case
		Token2: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/169e7244-7960-42dd-8a10-c7d24e97b24b",
			"https://open.feishu.cn/open-apis/bot/v2/hook/c0fa4970-8b25-4e24-a1c0-63c8b789387f",
			"https://open.feishu.cn/open-apis/bot/v2/hook/1b990120-101c-4020-b414-89fbc12dfcba",
		}, //富豪
		Token3: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/4c918cd9-b26e-4d06-bd20-2176b82ef68a",
			"https://open.feishu.cn/open-apis/bot/v2/hook/e807c4b8-7553-40b0-8537-aba766c07867",
			"https://open.feishu.cn/open-apis/bot/v2/hook/efa87b95-f208-4f75-b64c-86afa67b229f",
		}, //145
		Token4: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/dd60c8d2-b981-48dc-b6c9-3bb375d2f2c8",
			"https://open.feishu.cn/open-apis/bot/v2/hook/b5b6800b-880f-4814-88ca-aa0a69f75220",
		}, //谜语人
		Token5: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/7beb44ec-bf13-48f9-830f-e53537e1160a",
			"https://open.feishu.cn/open-apis/bot/v2/hook/a249ce33-5251-4f2f-bfbc-502cff6c6ea8",
			"https://open.feishu.cn/open-apis/bot/v2/hook/3cc9cb40-8d69-4cc3-bf20-381e353d609c",
		}, //柴郡猫
		Token6: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/c117e413-0b7b-49a7-92da-f6f9b5f1e547",
			"https://open.feishu.cn/open-apis/bot/v2/hook/5e1079f8-a73d-4f79-b6cc-021fc657df6d",
			"https://open.feishu.cn/open-apis/bot/v2/hook/06c76d39-fb61-44c2-a375-905d383e667b",
		}, //配置
		Token7: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/e301d22f-94fc-490d-a344-7f105acf0d7a",
		}, //陆逊
		Token8: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/54a59314-d0eb-4360-b6e5-05ec6467602f",
		}, //157胡莹
	}
)

func FeiShuUrl(text string, token string) {
	payload := &FeishuReq{
		MsgType: "text",
		Content: fmt.Sprintf("{\"text\":\"%s\"}", text),
	}
	jsonBytes, _ := json.Marshal(payload)
	if _, ok := tokenUrlMap[token]; ok {
		for _, v := range tokenUrlMap[token] {
			_, err := cli.Post(v, nil, jsonBytes)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
