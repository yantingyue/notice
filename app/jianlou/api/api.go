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

var SecondIdMap = make(map[uint64]struct{})

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
				"pageSize":            PageSize,
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
			sellInfo := sellInfo
			if _, ok := SecondIdMap[sellInfo.SecondId]; ok {
				if len(SecondIdMap) >= 20 {
					SecondIdMap = make(map[uint64]struct{})
				}
				continue
			}
			switch PayType {
			case 1:
				go func() {
					CreateOrderWallet(ctx, sellInfo.SecondId)
				}()
			case 2:
				go func() {
					CreateOrderKft(ctx, sellInfo.SecondId)
				}()
			}
			SecondIdMap[sellInfo.SecondId] = struct{}{}
			fmt.Println(SecondIdMap)
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
			"pay_pwd":  Pwd,
			"order_id": createOrderResp.Data.OrderId,
		}
		payResp := request(BuyToken, payReq, Urls[4])
		paySuccess := PayOrderResp{}
		json.Unmarshal(payResp, &paySuccess)
		if paySuccess.Code == 0 {
			cli.RedisClient.Incr(ctx, cast.ToString(ProductId))
			go func() {
				FeiShuUrl()
			}()
		}
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
				"pay_pwd":      Pwd,
			}
			payOrderResp := request(BuyToken, payReq, Urls[3])
			paySuccess := PayOrderResp{}
			json.Unmarshal(payOrderResp, &paySuccess)
			if paySuccess.Code == 0 {
				cli.RedisClient.Incr(ctx, cast.ToString(ProductId))
				go func() {
					FeiShuUrl()
				}()
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
