package feishu_notice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"log"
	"net/http"
	"notice/internal/cli"
	"strings"
)

func NiceNotice(token string) {
	ReqList(token)
}

func ReqList(token string) {
	client := &http.Client{}
	var data = strings.NewReader(`nice-sign-v1://b2d0044c2b7c7a541566d30265d08d4f:f069e3552d0a6d1e/{"nextkey":"","token":"MwGcQTsTGXTmPwDmPvTnPXiNQkwc-Ar6","search_key":"","type":"storage2goat","tag":"all","messageid":""}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/Sneakerstorage/list?a_x=-0.228592&a_y=-0.060287&a_z=-0.977768&abroad=no&appv=5.9.28.21&ch=AppStore_6.9.28.21&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=0.005967&g_y=-0.000289&g_z=-0.003649&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=192.168.2.1&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&seid=3656ea5e3ea6bbdc22de96f3fe07dc17&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&src=me&sw=430.000000&token=MwGcQTsTGXTmPwDmPvTnPXiNQkwc-Ar6&ts=1701535108409", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Cookie", "acw_tc=0bd17c5e17015343503837594e77b5fd988c91994a16c34dce2003ca837ea3")
	req.Header.Set("accept", "*/*")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("user-agent", "KKShopping/5.9.28 (iPhone 14 Pro Max; iOS 17.1.1; Scale/3.00)")
	req.Header.Set("accept-language", "zh-Hans-CN;q=1")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(bodyText) == 0 {
		return
	}
	goodResp := GoodsResp{}
	json.Unmarshal(bodyText, &goodResp)
	fmt.Printf("%s\n", bodyText)

	if goodResp.Code == 0 && len(goodResp.Data.List) > 0 {
		var text string
		ctx := context.Background()
		resultMap := make(map[string]struct{})
		for _, v := range goodResp.Data.List {
			resultMap[v.GoodsInfo.Id] = struct{}{}
			noticeKey := fmt.Sprintf("%s:%s", token, v.GoodsInfo.Id)
			cacheAll, err := cli.RedisClient.HGetAll(ctx, noticeKey).Result()
			if err != nil && err != redis.Nil {
				return
			}
			var onsale string
			for _, j := range v.OfferList {
				onsale = j.OnSale
			}
			cli.RedisClient.HMSet(ctx, noticeKey, map[string]interface{}{
				"id":         v.GoodsInfo.Id,
				"c":          v.Total.Num,
				"is_on_sale": onsale,
			})

			if len(cacheAll) == 0 {
				text = fmt.Sprintf("购买了《%s》总量%s个", v.GoodsInfo.Name, v.Total.Num)
				FeiShuUrlNice(text, token)
				continue
			} else {
				if v.Total.Num < cacheAll["c"] {
					text = fmt.Sprintf("《%s》卖出了1个了,剩余%s个", v.GoodsInfo.Name, v.Total.Num)
					FeiShuUrlNice(text, token)
				} else if v.Total.Num > cacheAll["c"] {
					text = fmt.Sprintf("购买了《%s》总量%s个", v.GoodsInfo.Name, v.Total.Num)
					FeiShuUrlNice(text, token)
				}
				if onsale != cacheAll["is_on_sale"] {
					switch onsale {
					case "-":
						text = fmt.Sprintf("《%s》暂无寄售,剩余%s个", v.GoodsInfo.Name, v.Total.Num)
					default:
						text = fmt.Sprintf("《%s》寄售中,剩余%s个", v.GoodsInfo.Name, v.Total.Num)
					}
					FeiShuUrlNice(text, token)
				}
			}

		}
		cacheJson, err := cli.RedisClient.Get(ctx, token).Result()
		if err != nil && err != redis.Nil {
			return
		}
		cli.RedisClient.Set(ctx, token, bodyText, 0)
		if cacheJson != "" {
			goodResp = GoodsResp{}
			json.Unmarshal([]byte(cacheJson), &goodResp)
			for _, v := range goodResp.Data.List {
				if _, ok := resultMap[v.GoodsInfo.Id]; !ok {
					text = fmt.Sprintf("《%s》卖光了", v.GoodsInfo.Name)
					cli.RedisClient.Del(ctx, fmt.Sprintf("%s:%s", token, v.GoodsInfo.Id))
					FeiShuUrlNice(text, token)
				}
			}
		}
	}
}

type GoodsResp struct {
	Code int32 `json:"code"`
	Data struct {
		List []struct {
			GoodsInfo struct {
				Id   string
				Name string
			} `json:"goods_info"`
			Total struct {
				Num  string
				Text string
			} `json:"total"`
			OfferList []struct {
				OnSale string `json:"on_sale"`
			} `json:"offer_list"`
		} `json:"list"`
	} `json:"data"`
}

type NiceNftList struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Result []struct {
			Id               uint64 `json:"id"`
			IsOnSale         uint32 `json:"is_on_sale"`
			C                uint32 `json:"c"`
			ProductTitle     string `json:"product_title"`
			ProductId        uint64 `json:"product_id"`
			NftProductSizeId uint64 `json:"nft_product_size_id"`
		} `json:"result"`
	} `json:"data"`
}
