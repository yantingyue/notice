package api

import (
	"encoding/json"
	"fmt"
	"log"
	"notice/internal/cli"
	"time"
)

var (
	list = LinkedList{}
)

func GetParam(token string, did string) map[string]interface{} {
	return map[string]interface{}{
		"token": token,
		"did":   did,
		"osn":   "IOS",
		"appv":  "5.9.28.21",
	}
}

func Begin() {
	go func() {
		for {
			list.Traverse(func(data interface{}) {
				log.Println(data)
				list.RemoveNode(data)
			})
		}
	}()
	for {
		for token, did := range TmpTokens {
			go func() {
				priceInfosV3(token, did)
			}()
			time.Sleep(time.Millisecond * 2000)
		}
	}
}

func priceInfosV3(token string, did string) {
	sign := Sign(fmt.Sprintf(`{"button_type":"purchase","id":"%d","sort":"common"}`, product_id), did, "")
	param := GetParam(token, did)
	resp := requestPost(param, sign, Urls[0])
	jsList := JsList{}
	json.Unmarshal(resp, &jsList)
	if jsList.Code == 0 && len(jsList.Data.TabList) > 0 && len(jsList.Data.TabList[0].List) > 0 {
		result := jsList.Data.TabList[0].List[0]
		if result.Stock > 0 {
			Config(SizeDetail{
				SizeId: result.SizeId,
				Stock:  result.Stock,
				Price:  result.Price,
			})
		}
	}
}

func Config(sizeDetail SizeDetail) {
	sign := Sign(fmt.Sprintf(`{"price":"%s","size_id":"%d","id":"%d","stock_id":"128"}`, sizeDetail.Price, sizeDetail.SizeId, product_id), BuyDid, "")
	param := GetParam(BuyToken, BuyDid)
	resp := requestPost(param, sign, Urls[1])
	configResp := ConfResp{}
	json.Unmarshal(resp, &configResp)
	if configResp.Code == 0 {
		prePub(configResp)
	}
}
func prePub(configResp ConfResp) {
	data := configResp.Data.StockInfo
	sign := Sign(fmt.Sprintf(`{"id":"%d","price":"%s","size_id":"%d","stock_id":"%d","unique_token":"%s","purchase_num":"1","need_storage":"yes"}`, product_id, data.Price, data.SizeId, data.StockId, configResp.Data.UniqueToken), BuyDid, "")
	param := GetParam(BuyToken, BuyDid)
	resp := requestPost(param, sign, Urls[2])
	prepubResp := PrepubResp{}
	json.Unmarshal(resp, &prepubResp)
	if prepubResp.Code == 0 {
		pub(configResp)
	}
	//fmt.Println(sign, "\n", param)
}
func pub(configResp ConfResp) {
	data := configResp.Data.StockInfo
	sign := Sign(fmt.Sprintf(`{"id":"%d","price":"%s","size_id":"%d","stock_id":"%d","unique_token":"%s","purchase_num":"1","need_storage":"yes"}`, product_id, data.Price, data.SizeId, data.StockId, configResp.Data.UniqueToken), BuyDid, "")
	param := GetParam(BuyToken, BuyDid)
	resp := requestPost(param, sign, Urls[3])
	prepubResp := PrepubResp{}
	json.Unmarshal(resp, &prepubResp)
	if prepubResp.Code == 0 {
		go func() {
			FeiShuUrl()
		}()
	}
}

func requestPost(param map[string]interface{}, body string, url string) (resp []byte) {
	req := cli.FastPostJson(url, param, []byte(body))
	resp, err := cli.FastResponse(req)
	if err != nil {
		log.Println(err)
	}
	list.AddNode(fmt.Sprintf("---%s---%s---%s", time.Now().Format("2006-01-02 15:04:05.000"), url, string(resp)))
	return
}
