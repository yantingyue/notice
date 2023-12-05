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
	var data = strings.NewReader(`nice-sign-v1://f45466d190f605ab7be46cb6f6d6ca37:6db61c9cd4c983f8/{"uid":"47537233","nextkey":"","comments_sort":"asc","mcc":"65535","isnewsession":1,"timestamp":0,"openudid":"b5280858f6c4a71527c227bc2bf9af19","mnc":"65535","token":"-jSRAt3QYLsjXm9WPkUTPIYPr9vBHECT","mark_read_sid":"","density":3,"ua":"Mozilla\/5.0 (iPhone; CPU iPhone OS 17_1_1 like Mac OS X) AppleWebKit\/605.1.15 (KHTML, like Gecko) Mobile\/15E148"}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/feed/userCollect?a_x=-0.001663&a_y=-0.044586&a_z=-0.992584&abroad=no&appv=5.9.29.24&ch=AppStore_6.9.29.24&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=-0.000010&g_y=-0.004267&g_z=-0.003943&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=114.114.114.114&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&pre_module_id=&seid=20c54c3716fab85872319d521ee9637e&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&sw=430.000000&token=-jSRAt3QYLsjXm9WPkUTPIYPr9vBHECT&tpid=user_profile&ts=1701738871605", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Cookie", "acw_tc=0bd17c0a17017388555616363ee55f83567c07fbef1130ecaf8dfdcc8e308e")
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
