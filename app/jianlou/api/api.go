package api

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.aiera.tech/cloud/comm/goroutine"
	"log"
	"notice/internal/cli"
	"notice/internal/util"
	"os"
	"time"
)

var (
	TmpTokens = []string{"24715fa709414f6eb364ffb6f8c13485"}
	Urls      = []string{
		"/aiera/ai_match_trading/nft_second/sell_product/list", //寄售列表
		"/aiera/ai_match_trading/nft_second/sell_order/pay",    //下单
		"/aiera/v2/hotdog/order/prepay",                        //预支付
		"/aiera/v2/hotdog/payment/kft/confirm",                 //支付

	}
)

const (
	TimeSpace        = 3000
	BuyNum           = 1
	BuyToken         = "8c131a620e0441b98fd0f4a3f6d946f4"
	ProductId        = 1019327
	NftProductSizeId = 1321
)

func Begin() {
	if len(TmpTokens) == 0 {
		return
	}
	ctx := context.Background()
	for {
		pageNum := 1
		for _, v := range TmpTokens {
			body := map[string]interface{}{
				"product_id":          ProductId,
				"nft_product_size_id": NftProductSizeId,
				"pageNumber":          pageNum,
				"pageSize":            BuyNum,
				"unlock":              1,
				"prop_pack":           0,
				"order_by":            "price",
			}
			goroutine.GoWithDefaultRecovery(ctx, func() {
				Grab(ctx, v, body)
			})
			time.Sleep(time.Millisecond * TimeSpace)
			i, _ := cli.RedisClient.Get(ctx, cast.ToString(ProductId)).Result()
			if cast.ToInt(i) == BuyNum {
				cli.RedisClient.Del(ctx, cast.ToString(ProductId))
				os.Exit(1)
			}
			pageNum++
		}
	}
}

func Grab(ctx context.Context, token string, body map[string]interface{}) {
	//查询寄售列表
	resp := request(token, body, Urls[0])
	sellList := SellListResp{}
	json.Unmarshal(resp, &sellList)
	if sellList.Code == 0 && len(sellList.Data.Res) > 0 {
		for _, sellInfo := range sellList.Data.Res {
			crOrderReq := map[string]interface{}{
				"operate_type":   "buy",
				"second_id":      sellInfo.SecondId,
				"user_coupon_id": 0,
			}
			//下单
			crOrderResp := requestOrder(BuyToken, crOrderReq, Urls[1])
			createOrderResp := CreateOrderResp{}
			json.Unmarshal(crOrderResp, &createOrderResp)
			if createOrderResp.Code == 0 && createOrderResp.Data.OrderId > 0 {
				//预支付
				prePayReq := map[string]interface{}{
					"pay_channel": 4,
					"order_id":    createOrderResp.Data.OrderId,
				}
				prePayResp := request(BuyToken, prePayReq, Urls[2])
				prePayOrderResp := PrePayOrderResp{}
				json.Unmarshal(prePayResp, &prePayOrderResp)
				if prePayOrderResp.Code == 0 && prePayOrderResp.Data.OrderNo != "" {
					//支付
					payReq := map[string]interface{}{
						"order_no":     prePayOrderResp.Data.OrderNo,
						"confirm_flag": "1",
						"pay_channel":  4,
						"pay_pwd":      "DVqBnIG8tFOmfbFp+tIXisluxkZDahm5Gk6MVvg4tY9td7tfjTvu5JiCDBmW39mUhgjY0z6zzlfj6Jc0/YDyaGLLB8n/wRXHoPRv6qlOyMleQw1iU5Y10MfF0jYylh2EJtiVd8VQWwOWgAuYmCIYUNqoy4IhjYxMs9Bj82l/rts=",
					}
					payOrderResp := request(BuyToken, payReq, Urls[3])
					paySuccess := PayOrderResp{}
					json.Unmarshal(payOrderResp, &paySuccess)
					if paySuccess.Code == 0 {
						cli.RedisClient.Incr(ctx, cast.ToString(ProductId))
					}
				}
			}
		}
	}
}

func request(token string, body map[string]interface{}, url string) (resp []byte) {
	header := util.GenerateHeader(token)
	jsonBytes, _ := json.Marshal(body)
	resp, _ = cli.Post(fmt.Sprintf("%s%s", Host, url), header, jsonBytes)
	log.Println("------------------", string(resp))
	if len(resp) == 0 {
		return resp
	}
	return resp
}
func requestOrder(token string, body map[string]interface{}, url string) (resp []byte) {
	header := util.GenerateCreateOrderHeader(token)
	jsonBytes, _ := json.Marshal(body)
	resp, _ = cli.Post(fmt.Sprintf("%s%s", Host, url), header, jsonBytes)
	log.Println(string(resp))
	if len(resp) == 0 {
		return resp
	}
	return resp
}

// 生成32位MD5 11
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

const (
	Host     = "https://api.aichaoliuapp.cn"
	timeOut  = 60 * time.Second
	version  = "31850"
	channel  = "010100"
	platform = "ios"
	appname  = "aiera.sneaker.snkrs.shoe"
	salt     = "5c33494d1b277902d1b78f98093f6fd4"
)
