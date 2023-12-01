package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"notice/internal/cli"
	"strings"
)

func B() {
	sign := "nice-sign-v1://c7caf4947a00b1bed5b17db04fbc5a95:cSnkcyuV9R8qWHdc/{\"button_type\":\"purchase\",\"id\":\"852749\",\"sort\":\"common\"}"
	RequeatList("R4KPSXOgbJTwJmeqpleWPiS8YeiBKuv_", sign, product_id)
	return
	for _, v := range TmpTokens {
		i := 0
		for {
			RequeatList(v, sign, product_id)
			i++
			if i == 3 {
				break
			}
			//time.Sleep(time.Millisecond * TimeSpace)
		}
	}
}
func ReqList() {
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
	fmt.Printf("%s\n", bodyText)
	goodResp := GoodsResp{}
	json.Unmarshal(bodyText, &goodResp)
	fmt.Println(goodResp)
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

func RequeatList(token string, sign string, id int) {
	client := &http.Client{}
	//d := fmt.Sprintf("%s{\"token\":\"%s\",\"id\":\"%d\",\"key\":\"\",\"button_type\":\"purchase\",\"sort\":\"common\",\"nextkey\":\"\"}", sign, token, id)
	var data = strings.NewReader(sign)
	//url := `https://api.oneniceapp.com/Sneakerpurchase/priceInfosV3?a_x=0.109711&a_y=-0.664032&a_z=-0.703842&abroad=no&appv=5.9.28.21&ch=AppStore_6.9.28.21&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=-0.001249&g_y=-0.004777&g_z=-0.001727&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=192.168.2.1&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&seid=837a9674a115bfffeaf0672ada417a62&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&src=goods_detail&sw=430.000000&token=` + fmt.Sprintf("%s&ts=1701273008134", token)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/Sneakerpurchase/priceInfosV3?a_x=-0.002670&a_y=-0.037247&a_z=-0.991714&abroad=no&appv=5.9.28.21&ch=AppStore_6.9.28.21&did=d5401cf612846e7cd15a2318039d67b8&dn=iPhone&dt=iPhone15%2C3&g_x=-0.255785&g_y=0.096176&g_z=-0.166534&geoacc=0&la=cn&lm=mobile&lp=-1.000000&n_bssid=&n_dns=114.114.114.114&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.1&seid=20e3bc4fc4c83303e14fca7478b2f582&sh=932.000000&sm_dt=2023112716423167b3d97f0dbec012f88c9d3b7e14ea5b01800fe81deb8019&src=goods_detail&sw=430.000000&token=R4KPSXOgbJTwJmeqpleWPiS8YeiBKuv_&ts=1701400608069", data)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "KKShopping/5.9.28 (iPhone 14 Pro Max; iOS 17.1.1; Scale/3.00)")
	req.Header.Set("Accept-Language", "zh-Hans-CN;q=1")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}

func request(param map[string]interface{}, body string, url string) (resp []byte) {
	//jsonBytes, _ := json.Marshal(body)
	req := cli.FastPostJson(fmt.Sprintf("%s%s", Host, url), param, []byte(body))
	resp, err := cli.FastResponse(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(resp))
	return
}
