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

		//"a158f96632784fc5ba3c177cd6d6abd1",
		//"695fea274d30407383d990b92972d679",
		//"b4a31298c3ea4b498549ec1de3dd6e6d",
		//"a7eb996385314faf943b15b516d63cdc",
		//"59db482aee9845cc9f7c509d4a8980e3",
		//"4fcd46b54b98402daed3194e9a4511f4",
		//"b2762df5926c4b74b1f85b79d93c61a1",
		//"f4c026dba1e44d219d242e74cfaa55c8",
		//"593dde88c5da4fe9a97a1e35082b710e",
		//"d7264142d433491a9ae4525eb0f17af9",
		//"33e5ef2108bc407ba2712f3f04a3b303",
		//"a2893d5fd20b4e0eb20e8d09bebdf7e0",
		//"943cffb4f6bb4c9d8b2135116ed69dc6",
		//"158b7ec94ac8451fba3550627b73f799",
		//"8891d0be4bb6420a9d29358145d2fa1c",
		//"892439ae58d24dee8039c7b72aa31f00",
		//"321b8ae2c06d4f1cbb2138c0c5283e3c",
		//"cf704474151c4f71aaf8f8617a69a3e9",
		//"5facb0baf26d4d31953355eea0233717",
		//"9de8086a8e99498d878464f666507b3b",
		//"1495b0f9dd834f3e8a94e3b0003cc4cf",
		//"bd704f7b68e340bf9d77149f587efcaf",
		//"9ad700051d4841858c3bd53287f524dc",
		//"cd9472aa03f34d4fb652119a41e4c5c1",
		//"814b62a945774ddd81d230ded1e3a441",
		//"5e5368b958094776846a576ee08c5ed3",
		//"f8908497623d47d3a61fbcae5d43db15",
		//"1bf6b13546d64798b20d0aa8c65b635e",
		//"1384a1c29aa743af84c95fbb7c1a64a7",
		//"c4403f0313774f88bdda746ad81526c6",
		//"a8facfe04dd24ce2a956f6e14b93c59b",
		//"152178f1a4b9474fb8e0b4c5748eb738",
		//"afaf2d63a8c2422eb1816d313d5ed053",
		//"a40d351bfbf9436b81a651d4b86a88ec",
		//"17072f5c9d744665b734a45f5aa6406a",
		//"5ca62416330a43ddb687be2bc5c8007b",
		//"838f0333d3964bf087282856497438e4",
		//"6976eb16a91943cda8546617cb2ad012",
		//"9300ab6d239e4aa8ac8da44e5e17f8c3",
		//"db7b7da64f0c40508533449b5c3c9abe",
		//"97e94393c8994a36aefbcb3b665a815d",
		//"d21c4c4aef674464baa82dc129ef4e48",
		//"0180a785b75741d3ab7d5db3c176a9a1",
		//"46571e016893494e9c0f5db1f58978bc",
		//"ba06831245034fb1a3ef1a36e76d7a4c",
		//"819ec12f791b43e1bd34a5602dc9bd39",
		//"35791413159442e5871c9cdb8a243cf2",
		//"c2491eb1892a4fc9b5c0ca5fb8a54a7b",
		//"b1634efcb94f44e8b8a789fa3ff3336a",
		//"2f935897476340c3ba2dd1076651890a",

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
	TimeSpace        = 140
	BuyNum           = 50
	BuyToken         = "8c131a620e0441b98fd0f4a3f6d946f4"
	ProductId        = 1020177
	NftProductSizeId = 2065

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
				"pageSize":            5,
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
