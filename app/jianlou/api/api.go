package api

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"notice/internal/cli"
	"notice/internal/util"
	"os"
	"time"
)

var (
	TmpTokens = []string{
		"88818b2970924542b5fff708ac483bea",
		"f304f2f75a0e43169209474cc5989f70",
		"a9921a79d89649fcb06c14ce051a8ea0",
		"516e43bfbc654380a428e1cf270a9106",
		"308dd5e23f2942fdbaf307bcbe8efd84",
		"328a156b373e4357b810a4b8ca2a072f",
		"7e5b07aad8ca43b59e9cbd338d0d5ff0",
		"bc14ff8966334b5fa812e8d9c3400349",
		"13d012047f0448909a0a24e7b983d38d",
		"5d5bdccd87fc40818ea9485b33827d10",

		"24715fa709414f6eb364ffb6f8c13485",
		"34a9ccce1e514809a7d9d327ed8ec1be",
		"9ffc543b213849c793d2298df41b51d1",
		"939293c845ab46d598547ce1cff16c8e",

		//"a3f938bfc7db4eaeb19ed3edbcd3fcdd",
		//"6894e604f80c45d7b576ddd87686b016",
		//"72634006ac8341ada26d5c1dd62ced9d",
		//"a9f7a3be2f88433bbd1069dd4e6af593",
		//"492add2bc2994d67ad4ca21d082108ca",
		//"798103ced0724357ba80b526ad75184d",
		//"6ec2fb4205bc429e84919dbe36bcd474",
		//"04f3fa80067a4273989f3edd41e58a41",
		//"ca40789d6cac40f893a8b040e21ba43e",
	}
	Urls = []string{
		"/aiera/ai_match_trading/nft_second/sell_product/list", //寄售列表
		"/aiera/ai_match_trading/nft_second/sell_order/pay",    //下单
		"/aiera/v2/hotdog/order/prepay",                        //预支付
		"/aiera/v2/hotdog/payment/kft/confirm",                 //支付

	}
)

const (
	TimeSpace        = 300
	BuyNum           = 20
	BuyToken         = "8c131a620e0441b98fd0f4a3f6d946f4"
	ProductId        = 1020328
	NftProductSizeId = 2183
)

func Begin() {
	if len(TmpTokens) == 0 {
		return
	}
	ctx := context.Background()
	for {
		for _, v := range TmpTokens {
			body := map[string]interface{}{
				"product_id":          ProductId,
				"nft_product_size_id": NftProductSizeId,
				"pageNumber":          1,
				"pageSize":            2,
				"unlock":              0,
				"prop_pack":           0,
				"order_by":            "price",
			}
			go func() {
				Grab(ctx, v, body)
			}()
			time.Sleep(time.Millisecond * TimeSpace)
			i, _ := cli.RedisClient.Get(ctx, cast.ToString(ProductId)).Result()
			if cast.ToInt(i) >= BuyNum {
				cli.RedisClient.Del(ctx, cast.ToString(ProductId))
				os.Exit(1)
			}
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
						i, _ := cli.RedisClient.Get(ctx, cast.ToString(ProductId)).Result()
						if cast.ToInt(i) >= BuyNum {
							cli.RedisClient.Del(ctx, cast.ToString(ProductId))
							os.Exit(1)
						}
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
	log.Println("------------------", string(resp), token)
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
