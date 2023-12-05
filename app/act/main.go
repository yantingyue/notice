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

var (
	rwMut     sync.RWMutex
	id        uint64
	orderInfo = make(map[string]ResponseData)
	tokens    = []string{
		"f50cec488ce34493b4062e38aa893255",
		"2908386d133d4eed921a966ff8ccd313",
		"907786050c784d61ab7a77d7f154a778",
		"b42c019b1a444aa7b5c1e4cde4872eac",
		"7c2e34b3841e46469aea90696783da18",
		"eed5639e01be4c38a0051c08eb05628c",
		"5af8406bb92140dcb68628e2f146cc7f",
		"268a2d3d72a24307969f77521eb54193",
		"d0ea5461f61c4a69af66c0f39565964a",
		"0a758cb918e04a34b711312389fafd46",
		"39848278bf124a7485afa70f526b75f5",
		"2aae06662b3a4d2e9d4e80e5642e3f86",
		"8250e7042f804473ab9e84f02d3d4461",
		"bceb1cf2538f4caa9c24ed6af67ecb0d",
		"6ae733ab91244e6e9e9b63048260ee8d",
		"d5eca1bad3e8449d9b73604e0498fc8a",
		"d96acb471a46480d95ab4ccfa095b9ec",
		"bcf61d6b83f949b1a988c9d356489e0e",
		"838bb66685004f4dba7232d869bac510",
		"5cf703f7f44348538607033a56d558b8",
		"28f8ee4621b748ed9aced94264219527",
		"809ecdf4d57640a2840b5df756f4aa93",
		"835d0c0f6786442e80bf7292e5d635ef",
		"e658ea949e4a4e24807ca89d5a8aed2f",
		"4ed5bca5267d4d43bb50c982c046df65",
		"b06cb28c748b49e89c417d0ae359984c",
		"6f7aab5009884ae9bf815abaa842b21b",
		"d7b0f2be88db4264a925d4c9e5dab97a",
		"e4e9b9dea3ff434a9b39e40cc916ec6a",
		"8cd96e65fb01411ea897a541ade82cab",
		"a222f7fe3f624e64a49ed47dcd15edb5",
		"4e9ff25de1024f77bcd2ceb318e87429",
		"c51b1a8bc1264928b66c25f27fbe21da",
		"8699af17be904e328f2b591e54ee5834",
		"77b415ce869c453385a79ea090e81de7",
		"d718f6d15d1e4843a98db32bb789cd4b",
		"7b74e0ca85c84a56bb1e60eedb8a7651",
		"7c96bccdc1d447e5a996e1d940aaf60d",
		"ed1566deb56c46d088f7bb9002674d73",
		"13e1deb79f7a4e08b6604bc1e0dacc66",
		"b1d42c0d15e14559a3b4c5caa9be8a35",
		"680593c81d9d49b8b46e5184ed41534b",
		"d13b65dc31484f7ead28ade8af6a2cbf",
		"d5bce03638b44555aab74728814ff994",
		"ab03784a7e2848b198b25acbac0974a4",
		"c99e91cb283a447fb609cf0dfb91fd05",
		"2e261988106b432eb268847fa6cc74f9",
		"dbe8dff903424eadb9cf8a6ec998bdf7",
		"e9ef0c2fb3d34ec6b010b6086930d917",
		"ed91376e29f54c4bb13db19a805a4ec2",
		"cb8f7f81e98f47ebb631090ee0e21266",
		"49e79be5c9c04151aeb7d507d781d1ed",
		"ed8a5489f2724ea89f7b67fb76407d85",
		"3d462c453c754581962cf18ea1dfe28f",
		"105aaf410081464c961899578cbf4205",
		"fb64c0a3824e4c5a8e2dc93f95e7a687",
		"ba80cd9a00a84b0ebff1e2078b364195",
		"c634fc85cca0459f91fd2203b0cf0caf",
		"410ecd2bd7ce440eb723dc9096b002c4",
		"e46e9aa7dabd43deb472a133d8f21fb7",
		"c15f6eccf6e44c7ca8d690c160b4f689",
		"0bcb42c6d0684ec5bd7e178538d7110b",
		"2c5f3e5c52ff48948d49ddced024daaf",
		"521f8e9a31ad4a70831ee1bf3b1301b4",
		"68226b0f26024a479f5dcebaca76841f",
		"356c930bafe94588bd79b22a4c9c1e20",
		"a4e5d5911b11413195972bfd424a5873",
		"9d3545612929414aa6b1340c149a09de",
		"f1d0479808124b33ad4ac7956191db89",
		"07a561c7e57a449b80e93adb64e80d1b",
		"8032b2a3932448d3a873d1c3e15f7260",
		"a24985c0b86a4777965512096e2307b2",
		"cbe67b1e2a7d4cd58ca9df8d5d65e57a",
		"98d0e23d80d84c85b7fd1607ef6f6338",
		"272a3544ac984975b0d86c98d2335c7a",
		"4174fd393dbb4a3589fbc8e1e9abbc1b",
		"e99e2b0d71db4eeab7bca0345de1b4e7",
		"f4efe7c981ee4eb69f6892b7c62ef1cf",
		"eacc05814fc9496e94e3b6e9a8d60b14",
		"67e29a07d9704c73ac4853796ddb4a07",
		"300dbdacbd274338a6efba57b89803d3",
		"3789e94fe896463ca48bf51911b6bcd1",
		"81108a47cf0344369eca9c2f629274a9",
		"df9f04d393a4478a8738253a1612624a",
		"215142fe12d64f00904cc0f1f018ae60",
		"189ffcab74754595822b0afe3f81eab2",
		"6904f46b1ef84f34940a108229a19c50",
		"08c9b4c2a47a4a138ccda0c7deddf509",
		"24ffc7a3d188414eb697ebdc130663b7",
		"a3abddc4ebba4d19b3174e64ca6284d0",
		"b24e3e4c5ca041eb9e96b58530fa771f",
		"e730b383ecd245c1ac25eb686663fb68",
		"cd771e7f2eba4b4a8bc3e4ef9288d624",
		"ce58ffac6c0c4c1491cca6356296d737",
		"a34bccd0a95f44df9a34625ed3157676",
		"ed5a9709b1a94812a68a28c131c3c96c",
		"dbfa81fc9e3d4f55ab341647051c0ecd",
		"eb2a21a5bd954d9cbe5afca904f4bd1f",
		"9e575818ded9423dbe7041840a13986f",
		"967c6d7775da409fa299f52d530ed682",
	}
	buyTokens = []string{
		"5dc91a4f8ebc4ffdb7de75c86433eb47", //yty
		"12186e8a5fd84bd78e0a9269181605e6", //pz
		//"f457f3597a04467bafe6172832ebe84d",  //zqq
	}
)

const (
	b     = 1   //1是分解 2是置换
	actId = 868 //活动id
)

func main() {
	go Begin()
	go Fj()
	select {}
}
func Begin() {
	for {
		if len(orderInfo) == 0 {
			for _, v := range buyTokens {
				orderRes := GetOrderInfo(actId, v)
				if len(orderRes.Data) > 0 {
					orderInfo[v] = orderRes
				}
			}
		}
		if len(orderInfo) > 0 {
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
				for _, token := range tokens {
					go func() {
						if FjDetail(actId, token) {
							//颜庭跃
							for k, v := range orderInfo {
								go func() {
									for j, item := range v.Data {
										if item.Type == "prop" {
											if ReplaceProp(actId, item.PropUserUuid, k) {
												fmt.Println(j)
												//rwMut.Lock() // 加写锁
												//if len(orderInfo[k].Data) > 1 {
												//	v.Data = orderInfo[k].Data[j+1:]
												//	orderInfo[k] = v
												//}
												//rwMut.Unlock() // 解写锁
											}
										} else {
											if Replace(actId, item.OrderID, k) {
												//fmt.Println(j)
												//rwMut.Lock() // 加写锁
												//if len(orderInfo[k].Data) > 1 {
												//	v.Data = orderInfo[k].Data[j+1:]
												//	orderInfo[k] = v
												//}
												//rwMut.Unlock() // 解写锁
											}
										}
									}
								}()
							}
						}
					}()
					time.Sleep(time.Millisecond * 10)
				}
			}
			//置换
			if b == 2 {
				for _, token := range tokens {
					go func() {
						if ReplaceDetail(actId, token) {
							//颜庭跃
							for k, v := range orderInfo {
								for j, item := range v.Data {
									fmt.Println(j)
									if item.Type == "prop" {
										if ReplaceProp(actId, item.PropUserUuid, k) {
											//rwMut.Lock() // 加写锁
											//if len(orderInfo[k].Data) > 1 {
											//	v.Data = orderInfo[k].Data[j+1:]
											//	orderInfo[k] = v
											//}
											//rwMut.Unlock() // 解写锁
										}
									} else {
										if Replace(actId, item.OrderID, k) {
											//rwMut.Lock() // 加写锁
											//if len(orderInfo[k].Data) > 1 {
											//	v.Data = orderInfo[k].Data[j+1:]
											//	orderInfo[k] = v
											//}
											//rwMut.Unlock() // 解写锁
										}
									}
								}
							}
						}
					}()
					time.Sleep(time.Millisecond * 15)
				}
			}
		}
	}()

}
func GetOrderInfo(id uint64, token string) (res ResponseData) {
	header := GenerateHeader1(token)
	body := map[string]interface{}{
		"replace_id": id,
		"pageNumber": 1,
		"pageSize":   50,
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
	log.Println(string(resp), token)
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
	if diffTime < 35 && diffTime > 0 {
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
	//resp, _ := Post("https://api.aichaoliuapp.cn/aiera/v2/hotdog/activity/displace/batch", header, jsonBytes)
	log.Println(orderId, string(resp), token)
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

	log.Println(propUUid, string(resp), token)
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
