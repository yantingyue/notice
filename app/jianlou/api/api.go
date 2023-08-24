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

		"cac29068ad1d45db88eb410c0ecdbafe",
		"4ece3f0db50148ecbb59923b34982f4f",
		"d1cd1ee21d7a40809ec3f655f720d744",
		"7f814785a5b5446fbb65f90564115d9b",
		"2a62e05347454fdf9d82c0b73f5eb5ea",
		"69a153491cc549c8a679aedf239e6d87",
		"8dfcb52533c949df9595613f25e82f85",
		"ccafbbd5a5f14485b7b3864e3a1acfb1",
		"265b6448e0824170a078dd877496db58",
		"f66e6f14ed5747d5b2895aafc34b85e1",
		"c06c766d13ac4ec4b659adba9997daa7",
		"2cf24880dce84da28deb5d48ca631103",
		"2398bbacf9ee439680899b28891d2f5c",
		"269defb055694470b8c9d2e5ddf24302",
		"10a7c1121fc748d2b651b53955dd3716",
		"c58495a0acb84e66adfc6cedd1ede3c0",
		"ff47ddd17e7d45bbad521a4bddf93669",
		"77fc1d99f36f443b9306f322760e42f2",
		"7cf891e406884ee6978d97d3cc2654d1",
		"918055f31a114417b0dc0a4d7605e94f",

		"f84e0824a37548bab559c36eacce2d1a",
		"0bbb0c2ea84847dd90657f1318b44f9d",
		"14eb87926a4543ea961646129da96ffd",
		"f20fab7739e54e5c9a0e16566def12a0",
		"f5159dc7230841228f449ee7c419c3f6",
		"a35c893822634079acc15a0bf76aae15",
		"9a0bb077a54b4afe9dbf2c6df660f0ea",
		"0dc668d4616b49fbb007d1b7d26ab628",
		"c44951f74ddd478f8cdf7057e0cbe1b8",
		"f3d2c9901cbb4c1597183395883fdc13",
		"baf55c06c23f4488a6c9eac5282487de",
		"d494ee0ab1aa48b0b4680b927384d92a",
		"802c95c2d2384497b592f9113f4fc954",
		"bff60ec7b1d84674acb54d2bf9e92616",
		"f755a66b57124b8587f7d7d7c37bc780",
		"af2fea66e9f14a3ab5da7c0062b255be",
		"d140e1ece17c4bd1b3936004dca4327b",
		"605ff5b45fb34ebfbe07234f6f4e5d40",
		"9e698e471529405e8924fd92f9533c9a",
		"392962c592374a5a8811223039bd918a",
		"04ebb5abec3144dd8da99be8ab8b41ca",
		"1d5dba394ba34b2fb1a67e561b90ef1f",
		"12416fe9ec524ec982ed68d838861280",
		"aedff138a86c4528882edd662523531e",
		"2f8960f930cb4f6ea68364c77e819c29",
		"abf9982b1fd34781a3df8ec7f150855c",
		"b7a0ae88626443beb57f5742b94dfe93",
		"580eada36921424886297ca14ddfb459",
		"4c01f5b22b7f4b8fba89de072cac6440",
		"d2871ee574414d7d9f129d60c79700ff",

		"24715fa709414f6eb364ffb6f8c13485",
		"34a9ccce1e514809a7d9d327ed8ec1be",
		"9ffc543b213849c793d2298df41b51d1",
		"939293c845ab46d598547ce1cff16c8e",
		"b369f2b8bc994c55909f457f47ff3a9d", //lzx

		"a3f938bfc7db4eaeb19ed3edbcd3fcdd",
		"72634006ac8341ada26d5c1dd62ced9d",
		"a9f7a3be2f88433bbd1069dd4e6af593",
		"492add2bc2994d67ad4ca21d082108ca",
		"798103ced0724357ba80b526ad75184d",
		"6ec2fb4205bc429e84919dbe36bcd474",
		"04f3fa80067a4273989f3edd41e58a41",
		"ca40789d6cac40f893a8b040e21ba43e",
		"0c2b35845f2a4c6bbc2dcc6ec4f22370",
		"7ec7bfd30c234a67b8063c0223efe6f9",
		"976660617b644b129d47fdb124e8c501",
		"8eca83e07c2c4d989fc709296b2dc2b0",
		"260f7ce6a2ff4e5884a5ca701b71f8e0",
		"b12a29d7ef674e1b98bbf33043af0d38",
		"84efec1d4aa84729991ba0300f0e9ac9",
	}
	Urls = []string{
		"/aiera/ai_match_trading/nft_second/sell_product/list",       //寄售列表
		"/aiera/ai_match_trading/nft_second/sell_order/pay",          //下单
		"/aiera/v2/hotdog/order/prepay",                              //快付通预支付
		"/aiera/v2/hotdog/payment/kft/confirm",                       //快付通支付
		"/aiera/ai_match_trading/nft_second/sell_order/wallet_order", //零钱支付
	}
)

const (
	TimeSpace = 140
	BuyNum    = 3
	BuyToken  = "8c131a620e0441b98fd0f4a3f6d946f4"
	//ProductId        = 1020177
	//NftProductSizeId = 2065

	ProductId        = 1019939
	NftProductSizeId = 1863

	//ProductId        = 1020338
	//NftProductSizeId = 2193
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
			//CreateOrderKft(ctx, sellInfo.SecondId)
			go func() {
				CreateOrderWallet(ctx, sellInfo.SecondId)
			}()
		}
	}
}
func CreateOrderWallet(ctx context.Context, secondId uint64) {
	crOrderReq := map[string]interface{}{
		"operate_type":   "buy",
		"second_id":      secondId,
		"user_coupon_id": 0,
	}
	//下单
	crOrderResp := requestOrder(BuyToken, crOrderReq, Urls[1])
	createOrderResp := CreateOrderResp{}
	json.Unmarshal(crOrderResp, &createOrderResp)
	if createOrderResp.Code == 0 && createOrderResp.Data.OrderId > 0 {
		//零钱支付
		payReq := map[string]interface{}{
			"pay_pwd":  "DVqBnIG8tFOmfbFp+tIXisluxkZDahm5Gk6MVvg4tY9td7tfjTvu5JiCDBmW39mUhgjY0z6zzlfj6Jc0/YDyaGLLB8n/wRXHoPRv6qlOyMleQw1iU5Y10MfF0jYylh2EJtiVd8VQWwOWgAuYmCIYUNqoy4IhjYxMs9Bj82l/rts=",
			"order_id": createOrderResp.Data.OrderId,
		}
		payResp := request(BuyToken, payReq, Urls[4])
		paySuccess := PayOrderResp{}
		json.Unmarshal(payResp, &paySuccess)
		//if paySuccess.Code == 0 {
		//	cli.RedisClient.Incr(ctx, cast.ToString(ProductId))
		//	i, _ := cli.RedisClient.Get(ctx, cast.ToString(ProductId)).Result()
		//	if cast.ToInt(i) >= BuyNum {
		//		cli.RedisClient.Del(ctx, cast.ToString(ProductId))
		//		os.Exit(1)
		//	}
		//}
	}
}

func CreateOrderKft(ctx context.Context, secondId uint64) {
	crOrderReq := map[string]interface{}{
		"operate_type":   "buy",
		"second_id":      secondId,
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
	log.Println(url, string(resp))
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
