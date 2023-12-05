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
	var data = strings.NewReader(`nice-sign-v1://1f459480a4a4d17481cf32b8efde309c:0aa003ab849069f6/{"uid":"47537233","nextkey":"","comments_sort":"asc","mcc":"65535","isnewsession":1,"timestamp":0,"openudid":"6fd27b323c1f7efdde8ae60166cec205","mnc":"65535","token":"gZro3dBSqaHSXwH_xeJ5PFJkiPn8sk26","mark_read_sid":"","density":3,"ua":"Mozilla\/5.0 (iPhone; CPU iPhone OS 17_1_2 like Mac OS X) AppleWebKit\/605.1.15 (KHTML, like Gecko) Mobile\/15E148"}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/feed/userCollect?a_x=-0.112137&a_y=-0.495331&a_z=-0.858841&abroad=no&appv=5.9.29.24&ch=AppStore_6.9.29.24&did=582b450952054561320f504965cf09a8&dn=iPhone&dt=iPhone15%2C3&g_x=-0.402609&g_y=0.665802&g_z=-0.068408&geoacc=0&la=cn&lm=weixin&lp=-1.000000&n_bssid=&n_dns=192.168.2.1&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.2&pre_module_id=&seid=e8bba31fcea19b4c8f0adf3857c0ced5&sh=932.000000&sm_dt=2023112723373870c182acedc58a6bad8ab30e3d1d1e94013bb877a8737acb&sw=430.000000&token=gZro3dBSqaHSXwH_xeJ5PFJkiPn8sk26&tpid=user_profile&ts=1701746145402", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "KKShopping/5.9.29 (iPhone 14 Pro Max; iOS 17.1.2; Scale/3.00)")
	req.Header.Set("Accept-Language", "zh-Hans-CN;q=1")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", "acw_tc=0bd17c5e17017457794617457e77b83e706399d6883c1d07eae335b86e9d86")
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
