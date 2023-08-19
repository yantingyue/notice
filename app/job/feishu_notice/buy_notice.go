package feishu_notice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"log"
	"notice/internal/cli"
	"time"
)

func MotorNotice(name string, userId uint64) {
	ctx := context.Background()
	token, _ := cli.RedisClient.Get(ctx, cast.ToString(userId)).Result()
	fmt.Println(token)
	if token == "" {
		tokens := requestToken(userId)
		if tokens == "" {
			text := fmt.Sprintf("获取token失败了")
			FeiShuUrl(text, userId)
			return
		}
		cli.RedisClient.Set(ctx, cast.ToString(userId), tokens, 0)
		token = tokens
	}
	handle(ctx, token, name, userId)
}

func handle(ctx context.Context, token string, name string, userId uint64) {
	pageSize := 50
	if name == "蜥蜴绝缘体" {
		pageSize = 100
	}
	body := map[string]interface{}{
		"pageSize":   pageSize,
		"pageNumber": 1,
		"page_type":  "art_create",
	}
	var text string
	fmt.Println("========================================", name, "=========================================")
	vv, res := request(token, body, userId)
	fmt.Println("========================================", name, "=============================================")
	fmt.Println("\n")
	if vv && len(res.Data.Result) > 0 {
		resultMap := make(map[uint64]struct{})
		for _, v := range res.Data.Result {
			resultMap[v.ProductId] = struct{}{}
			if v.ProductId == 1019723 || v.ProductId == 1019287 || v.ProductId == 1019119 || v.ProductId == 1019555 {
				continue
			}
			noticeKey := fmt.Sprintf("%s:%d", name, v.ProductId)
			cacheAll, err := cli.RedisClient.HGetAll(ctx, noticeKey).Result()
			if err != nil && err != redis.Nil {
				return
			}
			cli.RedisClient.HMSet(ctx, noticeKey, map[string]interface{}{
				"id":         v.Id,
				"c":          v.C,
				"is_on_sale": v.IsOnSale,
				"product_id": v.ProductId,
			})
			if len(cacheAll) == 0 {
				text = fmt.Sprintf("购买了《%s》总量%d个", v.ProductTitle, v.C)
				FeiShuUrl(text, userId)
				continue
			} else {
				if v.C < cast.ToUint32(cacheAll["c"]) {
					text = fmt.Sprintf("《%s》卖出了1个了,剩余%d个", v.ProductTitle, v.C)
					FeiShuUrl(text, userId)
				} else if v.C > cast.ToUint32(cacheAll["c"]) {
					text = fmt.Sprintf("购买了《%s》总量%d个", v.ProductTitle, v.C)
					FeiShuUrl(text, userId)
				}
				if v.IsOnSale != cast.ToUint32(cacheAll["is_on_sale"]) {
					switch v.IsOnSale {
					case 0:
						text = fmt.Sprintf("《%s》暂无寄售,剩余%d个", v.ProductTitle, v.C)
					case 1:
						text = fmt.Sprintf("《%s》寄售中,剩余%d个", v.ProductTitle, v.C)
					}
					FeiShuUrl(text, userId)
				}
			}
		}

		cacheJson, err := cli.RedisClient.Get(ctx, name).Result()
		if err != nil && err != redis.Nil {
			return
		}
		b, _ := json.Marshal(res)
		cli.RedisClient.Set(ctx, name, string(b), 0)
		if cacheJson != "" {
			nftList := &NftList{}
			json.Unmarshal([]byte(cacheJson), &nftList)
			for _, v := range nftList.Data.Result {
				if _, ok := resultMap[v.ProductId]; !ok {
					text = fmt.Sprintf("《%s》卖光了", v.ProductTitle)
					FeiShuUrl(text, userId)
				}
			}
		}
	}
}

type NftList struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Result []struct {
			Id           uint64 `json:"id"`
			IsOnSale     uint32 `json:"is_on_sale"`
			C            uint32 `json:"c"`
			ProductTitle string `json:"product_title"`
			ProductId    uint64 `json:"product_id"`
		} `json:"result"`
	} `json:"data"`
}

func request(token string, body map[string]interface{}, userId uint64) (bool, NftList) {
	res := NftList{}
	header := GenerateHeader1(token)
	jsonBytes, _ := json.Marshal(body)
	resp, _ := cli.Post(fmt.Sprintf("%s%s", Host, "/aiera/ai_match_trading/nft/order/list2"), header, jsonBytes)
	log.Println(string(resp))
	if len(resp) == 0 {
		return false, res
	}
	json.Unmarshal(resp, &res)
	if res.Code == 401 {
		cli.RedisClient.Del(context.Background(), cast.ToString(userId))
		go func() {
			FeiShuUrl("账号失效，请尽快处理", userId)
		}()
	}
	if res.Code == 0 && res.Msg == "success" {
		return true, res
	}
	return false, res
}

func GenerateHeader1(token string) map[string]string {
	timestamp := cast.ToString(time.Now().UnixMilli())
	return map[string]string{
		"token":     token,
		"version":   version,
		"channel":   channel,
		"platform":  platform,
		"appname":   appname,
		"timestamp": timestamp,
		"sign":      MD5(timestamp + salt),
	}
}

// 生成32位MD5 11
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func requestToken(userId uint64) string {
	head := make(map[string]string)
	auth, _ := cli.RedisClient.Get(context.Background(), "auth").Result()
	if auth == "" {
		return ""
	}
	head["Authorization"] = auth
	resp, _ := cli.Post(fmt.Sprintf("%s%d", "https://cmdb-api.aichaoliuapp.cn/api/get_prod_redis_value?key=", userId), head, nil)
	log.Println(string(resp))
	if len(resp) == 0 {
		return ""
	}
	resSt := RespToken{}
	json.Unmarshal(resp, &resSt)
	if resSt.Code != 200 {
		return ""
	}
	return resSt.Data.Value
}

type RespToken struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token string `json:"token"`
		Value string `json:"value"`
	} `json:"data"`
}
