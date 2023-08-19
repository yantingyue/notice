package feishu_notice

import (
	"encoding/json"
	"fmt"
	"log"
	"notice/internal/cli"
	"time"
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
	tokenUrlMap = map[int][]string{
		//Token1: []string{
		//	"https://open.feishu.cn/open-apis/bot/v2/hook/06a68041-bcdf-4f50-86d0-dd5a9a6012ed",
		//	"https://open.feishu.cn/open-apis/bot/v2/hook/186251cb-e289-4f5e-a06c-a708706c8eb9",
		//	"https://open.feishu.cn/open-apis/bot/v2/hook/1d35df52-2e54-45db-89f7-3a7383dd3958",
		//}, //case
		UserId2: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/169e7244-7960-42dd-8a10-c7d24e97b24b",
			"https://open.feishu.cn/open-apis/bot/v2/hook/c0fa4970-8b25-4e24-a1c0-63c8b789387f",
			"https://open.feishu.cn/open-apis/bot/v2/hook/1b990120-101c-4020-b414-89fbc12dfcba",
		}, //富豪
		UserId3: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/4c918cd9-b26e-4d06-bd20-2176b82ef68a",
			"https://open.feishu.cn/open-apis/bot/v2/hook/e807c4b8-7553-40b0-8537-aba766c07867",
			"https://open.feishu.cn/open-apis/bot/v2/hook/efa87b95-f208-4f75-b64c-86afa67b229f",
		}, //145
		UserId4: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/dd60c8d2-b981-48dc-b6c9-3bb375d2f2c8",
			"https://open.feishu.cn/open-apis/bot/v2/hook/b5b6800b-880f-4814-88ca-aa0a69f75220",
			"https://open.feishu.cn/open-apis/bot/v2/hook/44923305-a834-442c-8cba-a4e56376d609",
		}, //谜语人
		UserId5: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/7beb44ec-bf13-48f9-830f-e53537e1160a",
			"https://open.feishu.cn/open-apis/bot/v2/hook/a249ce33-5251-4f2f-bfbc-502cff6c6ea8",
			"https://open.feishu.cn/open-apis/bot/v2/hook/f71eb637-b7a8-4d8c-bfa8-1ac896058904",
		}, //柴郡猫
		UserId6: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/c117e413-0b7b-49a7-92da-f6f9b5f1e547",
			"https://open.feishu.cn/open-apis/bot/v2/hook/5e1079f8-a73d-4f79-b6cc-021fc657df6d",
			"https://open.feishu.cn/open-apis/bot/v2/hook/06c76d39-fb61-44c2-a375-905d383e667b",
		}, //配置
		UserId7: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/54a59314-d0eb-4360-b6e5-05ec6467602f",
		}, //157胡莹
		UserId8: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/1d58b5f2-cb0d-424d-b9b5-937c8dfc4270",
			"https://open.feishu.cn/open-apis/bot/v2/hook/fa5b77c2-2adb-4f57-9405-987157a26fc3",
		}, //门童
		UserId9: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/cace9648-9897-45ac-8ba3-67f83f4bc401",
			"https://open.feishu.cn/open-apis/bot/v2/hook/3b834a23-c6b9-4fc0-b322-d077b7700fc2",
		}, //富豪小号
		UserId10: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/cdd921a5-9da6-4c27-b7cc-211b456c4676",
			"https://open.feishu.cn/open-apis/bot/v2/hook/9bfa3673-e04c-45a3-ab8f-102b7c963fb1",
		}, //蜥蜴绝缘体
		UserId11: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/6298f961-3779-4a69-b1b5-d7fbd4452d6d",
			"https://open.feishu.cn/open-apis/bot/v2/hook/8e4989c5-7788-404c-8d1f-8e7fc96fa16a",
		}, //法典人
	}
)

func FeiShuUrl(text string, userId uint64) {
	payload := &FeishuReq{
		MsgType: "text",
		Content: fmt.Sprintf("{\"text\":\"%s\"}", text),
	}
	jsonBytes, _ := json.Marshal(payload)
	if _, ok := tokenUrlMap[int(userId)]; ok {
		for k, v := range tokenUrlMap[int(userId)] {
			if k > 0 && k < 2 {
				time.Sleep(time.Second * 3)
			}
			//go func() {
			_, err := cli.Post(v, nil, jsonBytes)
			if err != nil {
				log.Println(err)
			}
			//}()

		}
	}
}
