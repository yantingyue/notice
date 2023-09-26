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
			"https://open.feishu.cn/open-apis/bot/v2/hook/b31959d4-736c-43b2-a6a6-a8bfb14902a9",
			"https://open.feishu.cn/open-apis/bot/v2/hook/bb071f80-f0b3-4f4d-bd77-d844face7bf2",
			"https://open.feishu.cn/open-apis/bot/v2/hook/56e97935-e1bf-43f7-ae1b-a7eea9dd43ad",
		}, //富豪
		UserId3: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/ca3d6a3e-aa33-4606-a4ae-aca0897f9610",
			"https://open.feishu.cn/open-apis/bot/v2/hook/06c27e25-256f-4e86-b607-fdfbbd542862",
			"https://open.feishu.cn/open-apis/bot/v2/hook/e4694bd6-0737-408c-b438-438879fae108",
		}, //145
		UserId4: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/75d0a0c1-9974-4459-91ad-242f1f14eba7",
			"https://open.feishu.cn/open-apis/bot/v2/hook/b60202ce-16c2-45c5-b6a6-d76f2d1599d3",
			"https://open.feishu.cn/open-apis/bot/v2/hook/e1b91772-461c-4de4-bdd6-37236ff68fe5",
		}, //谜语人
		UserId5: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/def27b09-8181-44b3-b438-d29ce2668342",
			"https://open.feishu.cn/open-apis/bot/v2/hook/c8164378-7cb0-4a14-89f0-0a64e85666b2",
			"https://open.feishu.cn/open-apis/bot/v2/hook/4259fbcb-9452-40f1-b466-76fe04e0c775",
		}, //柴郡猫
		UserId6: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/2ffb3123-de49-4707-b0d3-6abbf77a44b5",
			"https://open.feishu.cn/open-apis/bot/v2/hook/82814345-a77e-4e54-ab73-bfdc469edc4f",
		}, //配置
		UserId7: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/e6c047f6-687a-4b9c-99dd-9882ab9183c6",
		}, //145胡莹
		UserId8: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/5df1bb38-5279-4ce6-83c6-1add87f11963",
			"https://open.feishu.cn/open-apis/bot/v2/hook/48bd771e-c4b0-4d35-8cc9-a885c7237b15",
		}, //门童
		UserId9: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/728f1f66-b232-4f09-a98b-b7b25aaf513b",
		}, //富豪小号
		UserId10: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/736ada58-1ce4-4b26-a6e1-c1b6ccc62667",
			"https://open.feishu.cn/open-apis/bot/v2/hook/630c830f-20f0-4885-9c0c-642245ac8e05",
			"https://open.feishu.cn/open-apis/bot/v2/hook/cd3fd2cd-6760-4264-8749-f81973ae03ca",
		}, //蜥蜴绝缘体
		UserId11: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/eb625e0f-a61f-4b1b-90ef-24e81952edda",
		}, //法典人
		UserId13: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/bc05ffbb-f5f6-4e4a-9c0d-e4c02d3a3154",
			"https://open.feishu.cn/open-apis/bot/v2/hook/8af7c7c8-e91d-420f-8a4e-3c7e3eaab705",
		}, //145韩新枝
		UserId14: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/485792e4-444a-4983-b01b-b35fec842e09",
		}, //陈凯歌
		UserId15: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/25db3d5c-2e26-427d-9adc-e9c41607202d",
		},
		UserId16: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/735993c3-0ee3-47ee-8eb1-da7eac69e208",
			"https://open.feishu.cn/open-apis/bot/v2/hook/06d4201e-9a56-4047-9cdf-0cad96e6f0ae",
		},
		UserId17: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/14910cd3-e362-4f92-a208-47ca43ceb5a1",
		},
		UserId18: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/d617e083-5412-481f-b9cb-dd53d63de8bc",
		},
		UserId19: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/ebc3317a-654a-4e13-9d22-408df8202352",
		},
		UserId20: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/309fed83-2159-432e-8727-9c965b639316",
		},
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
				time.Sleep(time.Second * 6)
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
