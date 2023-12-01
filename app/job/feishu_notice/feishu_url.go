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
			"https://open.feishu.cn/open-apis/bot/v2/hook/e1b4f2ef-7052-48cb-85c3-0e225bec83d1",
		}, //富豪小号
		UserId10: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/75b64def-18db-4f70-8cbc-9a770bec45fc",
		}, //蜥蜴绝缘体
		UserId11: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/9f4cf5e0-581d-4358-9722-8484e88346f1",
			"https://open.feishu.cn/open-apis/bot/v2/hook/7eb2eefc-b889-4d50-b60a-c27ff4efd1c2",
		}, //法典人
		UserId13: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/bc05ffbb-f5f6-4e4a-9c0d-e4c02d3a3154",
			"https://open.feishu.cn/open-apis/bot/v2/hook/8af7c7c8-e91d-420f-8a4e-3c7e3eaab705",
		}, //145韩新枝
		UserId14: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/485792e4-444a-4983-b01b-b35fec842e09",
		}, //陈凯歌
		UserId15: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/00c90c98-49a3-4590-a0d2-007f0f70080e",
			"https://open.feishu.cn/open-apis/bot/v2/hook/c3aaf54d-f473-498b-9904-bf9fee950713",
		},
		UserId16: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/17086938-e717-4306-a350-29cb1abd1f46",
			"https://open.feishu.cn/open-apis/bot/v2/hook/735993c3-0ee3-47ee-8eb1-da7eac69e208",
			"https://open.feishu.cn/open-apis/bot/v2/hook/06d4201e-9a56-4047-9cdf-0cad96e6f0ae",
		},
		UserId17: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/14910cd3-e362-4f92-a208-47ca43ceb5a1",
		},
		UserId18: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/124d8a3f-e69a-43f9-a649-b83bb0671b2d",
		},
		UserId19: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/ebc3317a-654a-4e13-9d22-408df8202352",
		},
		UserId20: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/309fed83-2159-432e-8727-9c965b639316",
		},
		UserId21: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/90dfeb55-3ff3-4209-b52a-aee0fc81c325",
		},
		UserId22: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/172781a5-10a0-4391-88f3-3cd4d72d4784",
		},
		UserId23: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/f66623ca-e15d-40cf-ac1d-f2b4fcf3755d",
		},
		UserId24: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/38005884-8a7f-44c5-94b8-ec22c3576339",
			"https://open.feishu.cn/open-apis/bot/v2/hook/a66aae5b-4a85-415d-af1a-befba0d609e3",
		},
		UserId25: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/7ea7971f-f8b0-49d7-b6bf-5e24fc319bc5",
			"https://open.feishu.cn/open-apis/bot/v2/hook/ec4a9ace-285a-4645-86f9-c8285e3fd750",
		},
		UserId26: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/679db88d-7e31-4990-837c-52e88c83861d",
		},
		UserId27: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/bb1a01e8-8dc2-4305-8aea-d10fea92be99",
		},
		UserId28: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/cf6d1481-d03c-4c2a-8502-b21f76cb1565",
			"https://open.feishu.cn/open-apis/bot/v2/hook/813b46d6-a16d-4518-9b10-6a74af3c7415",
		},
		UserId29: []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/ea0c24ab-5567-485e-a054-986176d6220c",
		},
	}

	nicetokenUrlMap = map[string][]string{
		"wEaOgGs2ulepxrsMlvimPoQSMxE3r3HO": []string{
			"https://open.feishu.cn/open-apis/bot/v2/hook/a3745359-72f0-47cf-8261-2b5e83af5ed6",
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

func FeiShuUrlNice(text string, token string) {
	payload := &FeishuReq{
		MsgType: "text",
		Content: fmt.Sprintf("{\"text\":\"%s\"}", text),
	}
	fmt.Println(payload)
	fmt.Println(token)
	jsonBytes, _ := json.Marshal(payload)
	if _, ok := nicetokenUrlMap[token]; ok {
		for _, v := range nicetokenUrlMap[token] {
			fmt.Println(v)
			_, err := cli.Post(v, nil, jsonBytes)
			if err != nil {
				log.Println(err)
			}

		}
	}
}
