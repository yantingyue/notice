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
	var data = strings.NewReader(`nice-sign-v1://ea5b7a1ccb314eeccb7663d8d17b4c99:90a02ffa3fd0091b/{"nextkey":"","token":"WXxr7urByMoVbnsdIsToPfTNwOl2W7k6","search_key":"","type":"storage2goat","tag":"all","messageid":""}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/Sneakerstorage/list?a_x=-0.005157&a_y=-0.037201&a_z=-0.994797&abroad=no&appv=5.9.28.21&ch=AppStore_6.9.28.21&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=0.006479&g_y=-0.002656&g_z=-0.007193&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=114.114.114.114&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&seid=54cf4faf244aee56214a9b64d0d7387e&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&src=me&sw=430.000000&token=WXxr7urByMoVbnsdIsToPfTNwOl2W7k6&ts=1701423017581", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Cookie", "acw_tc=0bd17c0a17014230065523868ee5396bd5100677e470f41ba2e1610528c46b; did=d5401cf612846e7cd15a2318039d67b8; id=54194172; lan=cn; name=01e2ca63fc687ef73a1961eee2d184ba; sign=a591c45565ed24568b72abb54edb6974; time=1701419821; token=WXxr7urByMoVbnsdIsToPfTNwOl2W7k6; uid=54194172; niceUser=%7B%22uid%22%3A%2254194172%22%2C%22rand%22%3A3451%2C%22expire%22%3A1702024621%2C%22sign%22%3A%226528c68ae8aa260f4740f6cdb496ab07%22%7D; lang=zh-cn; nuid=rBAAC2VpiYm08zLi8CaLAg==")
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
					text = fmt.Sprintf("购买了《%s》卖出了1个了,剩余%s个", v.GoodsInfo.Name, v.Total.Num)
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
		//cacheJson, err := cli.RedisClient.Get(ctx, token).Result()
		//if err != nil && err != redis.Nil {
		//	return
		//}
		//b, _ := json.Marshal(res)
		//cli.RedisClient.Set(ctx, name, string(b), 0)
		//if cacheJson != "" {
		//	nftList := &NftList{}
		//	json.Unmarshal([]byte(cacheJson), &nftList)
		//	for _, v := range nftList.Data.Result {
		//		if _, ok := resultMap[v.ProductId]; !ok {
		//			text = fmt.Sprintf("《%s》卖光了", v.ProductTitle)
		//			cli.RedisClient.Del(ctx, fmt.Sprintf("%s:%d:%d", name, v.ProductId, v.NftProductSizeId))
		//			FeiShuUrl(text, userId)
		//		}
		//	}
		//}
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
