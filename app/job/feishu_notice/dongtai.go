package feishu_notice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"notice/internal/cli"
	"strings"
	"time"
)

func DTList(token string) {
	client := &http.Client{}
	var data = strings.NewReader(`nice-sign-v1://1dd3c6d9de9a0f5f9b66f9c33cf982de:3b38ac493349136c/{"uid":"47537233","token":"6iFu2817SVTNpYXOPlsZP4jXsYINwUvS"}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/feed/pubListV2?a_x=-0.004379&a_y=-0.038559&a_z=-0.994217&abroad=no&appv=5.9.29.24&ch=AppStore_6.9.29.24&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=0.000254&g_y=-0.004688&g_z=-0.002148&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=114.114.114.114&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&pre_module_id=&seid=a4469022a3a79a4d437d80cd850a3542&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&sw=430.000000&token=6iFu2817SVTNpYXOPlsZP4jXsYINwUvS&tpid=user_profile&ts=1701740712532", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Cookie", "did=d5401cf612846e7cd15a2318039d67b8; id=54194172; lan=cn; name=%E7%83%AD%E7%8B%97%E5%86%85%E5%B9%95%E5%93%A5; sign=23e263e2893dd62c35474e800d36f6ea; time=1701740508; token=6iFu2817SVTNpYXOPlsZP4jXsYINwUvS; uid=54194172; nuid=rBAAF2VufU2Db1dlQVtcAg==; acw_tc=0bd17c0a17017398533415818ee53198b46074a792acee776e233f52ff29f1")
	req.Header.Set("accept", "*/*")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("user-agent", "KKShopping/5.9.29 (iPhone 14 Pro Max; iOS 17.1.1; Scale/3.00)")
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
	fmt.Printf("%s\n", bodyText)
	goodResp := DtResp{}
	fmt.Println(goodResp)
	json.Unmarshal(bodyText, &goodResp)
	if goodResp.Code == 0 && len(goodResp.Data.Timeline) > 0 {
		ctx := context.Background()
		for _, v := range goodResp.Data.Timeline {
			i, _ := cli.RedisClient.Get(ctx, cast.ToString(v.TradeDynamic.Id)).Result()
			if i != "" {
				continue
			}
			cli.RedisClient.Set(ctx, cast.ToString(v.TradeDynamic.Id), "1", time.Second*86400*10)
			FeiShuUrlNice(fmt.Sprintf("个人动态购买了《%s》价格%s元", v.TradeDynamic.SizeLabel, v.TradeDynamic.Price), "testFuhao")
		}
	}
}

type DtResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Timeline []struct {
			TradeDynamic struct {
				Id        uint64 `json:"id"`
				SizeLabel string `json:"size_label"`
				Price     string `json:"price"`
			} `json:"trade_dynamic"`
		} `json:"timeline"`
	} `json:"data"`
}