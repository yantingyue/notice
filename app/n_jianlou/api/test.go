package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Test() {
	client := &http.Client{}
	var data = strings.NewReader(`nice-sign-v1://f455ec7b3f2cc9c1a38e439251f9503a:2ff9e9a1de889100/{"password":"LwaQkBLnfzhC58LovoUmF+qjmtiBm\/PmMqIdm2LBaL73M0mIAXJZOzdW13uIFc3ux7LbNpDIc3Pr8FGas0I84VggpxubK1OCCYA2R6fVCM5iW4lRq8IMD76qaUZLM+bOP6ntLrJ+bgcVYSR3GLDu1XiDG8OJPY0AzWJ6ri2AwWg=","country":"1","mobile":"19121375684","platform":"mobile"}`)
	req, err := http.NewRequest("POST", "https://api.oneniceapp.com/account/login?a_x=-0.029160&a_y=-0.116150&a_z=-0.980576&abroad=no&appv=5.9.29.24&ch=AppStore_6.9.29.24&did=582b450952054561320f504965cf09a8&dn=iPhone&dt=iPhone15%2C3&g_x=-0.000088&g_y=-0.002857&g_z=-0.002254&geoacc=0&la=cn&lm=sina&lp=-1.000000&n_bssid=&n_dns=192.168.2.1&n_ssid=&net=0-0-wifi&osn=iOS&osv=17.1.2&seid=5fef33dbd0bee61c6a649fd0bf50f067&sh=932.000000&sm_dt=2023112723373870c182acedc58a6bad8ab30e3d1d1e94013bb877a8737acb&sw=430.000000&tpid=login&ts=1701829050360", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.oneniceapp.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "KKShopping/5.9.29 (iPhone 14 Pro Max; iOS 17.1.2; Scale/3.00)")
	req.Header.Set("Accept-Language", "zh-Hans-CN;q=1")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", "acw_tc=0bd17c5e17018290325455584e777a7ad70ae0c3ed6f38ad8776646c428e56")
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

func NiceSign() {

	sign := Sign(`{"uid":"47537233","nextkey":"","comments_sort":"asc","mcc":"65535","isnewsession":1,"timestamp":0,"openudid":"6fd27b323c1f7efdde8ae60166cec205","mnc":"65535","token":"h4alcZpramfMSt9JJesTPI87KowtJ0DG","mark_read_sid":"","density":3,"ua":"Mozilla\/5.0 (iPhone; CPU iPhone OS 17_1_2 like Mac OS X) AppleWebKit\/605.1.15 (KHTML, like Gecko) Mobile\/15E148"}`, "582b450952054561320f504965cf09a8", "")
	fmt.Println(sign)
}
