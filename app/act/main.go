package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/valyala/fasthttp"
)

type ReplaceResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OnSaleStatus       uint32 `json:"on_sale_status"`
		CurrentMilliTime   int64  `json:"current_milli_time"`
		StartTimeTimestamp int64  `json:"start_time_timestamp"`
		EndTimestamp       int64  `json:"end_time_timestamp"`
	} `json:"data"`
}
type ReplaceTimeResp struct {
	Code             int32  `json:"code"`
	Msg              string `json:"msg"`
	CurrentMilliTime int64  `json:"current_milli_time"`
}

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		OrderID      uint64 `json:"order_id"`
		Title        string `json:"title"`
		Picture      string `json:"picture"`
		UpdateTime   string `json:"update_time"`
		SubTitle     string `json:"sub_title"`
		PropUserUuid string `json:"prop_user_uuid"`
		Type         string `json:"type"`
		Weight       int    `json:"weight"`
	} `json:"data"`
}

const (
	b               = 1                                  //1是分解 2是置换
	actId           = 619                                //活动id
	thread          = 2                                  //并发数
	tokenCommon     = "da01634063c446659313a5a1e013f86c" //勿删
	tokenYanTingYue = "da01634063c446659313a5a1e013f86c" //颜庭跃
)

var (
	id        uint64
	orderInfo ResponseData
)

func main() {
	go Begin()
	go Fj()
	select {}
}
func Begin() {
	for {
		if len(orderInfo.Data) == 0 {
			orderInfo = GetOrderInfo(actId, tokenYanTingYue)
		}
		if len(orderInfo.Data) > 0 {
			break
		}
		fmt.Println("材料不足")
		time.Sleep(time.Millisecond * 1000)
	}

}

func Fj() {
	go func() {
		for {
			//分解
			if b == 1 {
				for i := 0; i < thread; i++ {
					go func() {
						if FjDetail(actId, tokenCommon) {
							//颜庭跃
							go func() {
								for k, v := range orderInfo.Data {
									if v.Type == "prop" {
										if ReplaceProp(actId, v.PropUserUuid, tokenYanTingYue) {
											orderInfo.Data = orderInfo.Data[k+1:]
										}
									} else {
										if Replace(actId, v.OrderID, tokenYanTingYue) {
											orderInfo.Data = orderInfo.Data[k+1:]
										}
									}
								}
							}()
						}
					}()
					time.Sleep(time.Millisecond * 30)
				}
			}
			//置换
			if b == 2 {
				for i := 0; i < thread; i++ {
					go func() {
						if ReplaceDetail(actId, tokenCommon) {
							for k, v := range orderInfo.Data {
								if Replace(actId, v.OrderID, tokenYanTingYue) {
									orderInfo.Data = orderInfo.Data[k+1:]
								}
							}
						}
					}()
				}
				time.Sleep(time.Millisecond * 20)
			}
		}
	}()

}
func GetOrderInfo(id uint64, token string) (res ResponseData) {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"replace_id": id,
		"pageNumber": 1,
		"pageSize":   5,
	}
	jsonBytes, _ := json.Marshal(body)
	resp, _ := Post("https://api.aichaoliuapp.cn/aiera/ai_match_trading/nft/combination/choice/material", header, jsonBytes)
	log.Println(string(resp))
	json.Unmarshal(resp, &res)
	return res
}

func FjDetail(id uint64, token string) bool {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"replace_id": id,
	}
	jsonBytes, _ := json.Marshal(body)
	resp, _ := Post("https://api.aichaoliuapp.cn/aiera/ai_match_trading/nft/replace/active/detail", header, jsonBytes)
	log.Println(string(resp))
	if len(resp) == 0 {
		return false
	}
	res := ReplaceResp{}
	json.Unmarshal(resp, &res)
	if res.Code == 0 && res.Msg == "success" && res.Data.OnSaleStatus == 1 {
		return true
	}
	return false
}
func ReplaceDetail(id uint64, token string) bool {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"replace_id": id,
	}
	jsonBytes, _ := json.Marshal(body)
	var (
		wg        sync.WaitGroup
		resDetail = ReplaceResp{}
		resTime   = ReplaceTimeResp{}
		resp1     []byte
		resp2     = make(map[string]interface{})
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp1, _ = Post("https://api.aichaoliuapp.cn/aiera/ai_match_trading/nft/replace/active/detail", header, jsonBytes)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp2, _ = Get("https://api.aichaoliuapp.cn/aiera/current/milli/time")
	}()
	wg.Wait()
	if len(resp1) == 0 || len(resp2) == 0 {
		return false
	}
	json.Unmarshal(resp1, &resDetail)
	g, _ := json.Marshal(resp2)
	json.Unmarshal(g, &resTime)
	//resDetail.Data.StartTimeTimestamp = 1689605160000
	diffTime := resDetail.Data.StartTimeTimestamp - resTime.CurrentMilliTime
	log.Println(resDetail.Data.StartTimeTimestamp, resTime.CurrentMilliTime, diffTime)
	if diffTime < 150 && diffTime > 0 {
		//time.Sleep(time.Millisecond * time.Duration(diffTime))
		log.Println(diffTime, resTime.CurrentMilliTime, resDetail.Data.StartTimeTimestamp, time.Now().UnixMilli())
		return true
	}
	if resDetail.Code == 0 && resDetail.Msg == "success" && resTime.CurrentMilliTime >= resDetail.Data.StartTimeTimestamp && resTime.CurrentMilliTime <= resDetail.Data.EndTimestamp {
		return true
	}
	return false
}
func Replace(id, orderId uint64, token string) bool {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"order_id":   orderId,
		"replace_id": id,
	}
	jsonBytes, _ := json.Marshal(body)
	resp, _ := Post("https://api.aichaoliuapp.cn/aiera/ai_match_trading/nft/replace/active/exchange", header, jsonBytes)

	log.Println(orderId, string(resp))
	if len(resp) == 0 {
		return false
	}
	res := ReplaceResp{}
	json.Unmarshal(resp, &res)
	if res.Code == 0 && res.Msg == "success" {
		return true
	}
	return false
}
func ReplaceProp(id int, propUUid string, token string) bool {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"replace_prop_uuid": propUUid,
		"replace_id":        id,
	}
	jsonBytes, _ := json.Marshal(body)
	resp, _ := Post("https://api.aichaoliuapp.cn/aiera/ai_match_trading/nft/replace/active/exchange", header, jsonBytes)

	log.Println(propUUid, string(resp))
	if len(resp) == 0 {
		return false
	}
	res := ReplaceResp{}
	json.Unmarshal(resp, &res)
	if res.Code == 0 && res.Msg == "success" {
		return true
	}
	return false
}

const (
	timeOut  = 60 * time.Second
	version  = "31850"
	channel  = "010100"
	platform = "ios"
	appname  = "aiera.sneaker.snkrs.shoe"
	salt     = "5c33494d1b277902d1b78f98093f6fd4"
)

func GenerateHeader1(token string) map[string]string {
	timestamp := cast.ToString(time.Now().UnixMilli())
	return map[string]string{
		"token":     token,
		"version":   version,
		"channel":   channel,
		"platform":  platform,
		"appname":   appname,
		"timestamp": timestamp,
		"sign":      MD5(timestamp + salt),
	}
}

func Get(host string) (m map[string]interface{}, err error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, host, timeOut)
	if err != nil {
		return
	}
	if statusCode != fasthttp.StatusOK {
		err = errors.New(fmt.Sprintf("request failed statusCode[%d]", statusCode))
		return
	}
	if body == nil {
		err = errors.New("response body is nil")
		return
	}
	m = make(map[string]interface{})
	if err = json.Unmarshal(body, &m); err != nil {
		return
	}
	return
}

func Post(host string, header map[string]string, payload []byte) (body []byte, err error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(host)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	req.SetBody(payload)

	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err = client.DoTimeout(req, resp, timeOut); err != nil {
		return
	}
	body = resp.Body()
	return
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
