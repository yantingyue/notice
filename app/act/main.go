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
		//111
		"12a89b1213064ddfa0bb931a8423442f",
		"014b06937c734734af2d463613b08f4d",
		"98d0b2f44723441eab78158bdd6e6823",
		"4897777728144572aaafa0237c314c3b",
		"585b9544be22413e97818080d559c8f7",
		"24b671d590184ed2ac078f4a38d8a09a",
		"fb5fc392a997402199d9ec2e266b4ac0",
		"4b678f79aed24c138154f12d58725725",
		"bd39e4eeed414e86bd7247364a98101b",
		"bd39e4eeed414e86bd7247364a98101b",
		"9a577e671b824bd0b022008c19c9ffb1",
		"116070fcd3334f70b65489124b861cdf",
		"0719c026b3e4478dbacfefd234a7c440",
		"d61e491e9e15470f8dfc6f3976ce8f03",
		"db35081be34049568c59f85c03e8b7e3",
		"60a94a911cab4d71bc8b76de5ef12985",
		"34c2c840286b4568b211fc3db983436a",
		"dc4cc7d29afb4449aafeba33e3848ccd",
		"c4ad7186154b45158b3830c6a9bec70e",
		"5a4a420df5e4439aa50188852b8b4269",
		"e957af3374ba4e45987b6c9ff5093f68",
		"d763bf52940e428d84c38b66b662bc76",
		"1215df75a1f947fcac2157e3e8022435",
		"c99070353dbb438d813b0831e0084925",
		"ef24cf8aeec642618c748f8cd8a0933f",
		"cad46ea5d45e4775ab2bc98e4a6598db",
		"5a08f30aaa2c4c809edd81f6ae43d6da",
		"906b299fb2af4c57ade6cfc0c5af6bd7",
		"11b25cd4e1af4746bf80933dae80af97",
		"2695aa83b0bb4b75b86e5888f4a14391",
		"73ce56294a5a48948e4411937e5f4f07",
		"3ea781f0db0244d1b9f9281f5f5f4837",
		"c7bec8d729354bd999ca7da7714dc4c8",
		"e992916c77424b3fa052997a25eb786d",
		"b8fd8481958b4e8cb95ddd96315625f3",
		"b86b7c24dec64aa99bffb1dac377cef4",
		"924a8f4719df47e9ac8c874a7c22616c",
		"af1643ac281f441e86d439c4a84b5e52",
		"af1643ac281f441e86d439c4a84b5e52",
		"5900ecc6dbea41d79e7accb422d12f5a",
		"0f27b5efc9ca473cb5b3a9c88fda6ba9",
		"d19743b66c1d469f88ee4ef9c20a56f8",
		"ce90615e69314179ab31ecb664ba2409",
		"c8ebd795269b4d159d9802be8b9a819f",
		"73bec595adc445e0bbc1808166ad7e13",
		"37be3d6496fc4d5abc83e6e255b986e3",
		"361d2aacb8b54f1588ef71feda3abbcf",
		"cd57351bec64423da2f84c45dc56b6d0",
		"7dadc3edac5d470ea505a08b3a1bb830",
		"7dadc3edac5d470ea505a08b3a1bb830",

		//22
		"66f6d63ec5e64496a350b54a2dc51ec7",
		"bfce241882e9486b9e3890719439a498",
		"53b0094d5e5548fbb8c048ec70d6dd8c",
		"cc253e7db6a44c23a950c587743ca022",
		"9ea78de591b34eb5b90ce8503e2ddd8e",
		"be0a23b24cc140ee9243a92992ab8fca",
		"8e4ec3af62f541458ac91071f14dbfa8",
		"d6d86e7052984edab19277a5ca4fd410",
		"bb63de3964004a9f97f48e891ef91965",
		"223b7d6820ea4f7189dd60dcd3dae569",
		"f0450cf165934bfb846c4e7aecdc9e0a",
		"a469e0f1feb840c5adf0f841aa4fc038",
		"8a00e4b7aeb347d896201e97d633df62",
		"6b6432122f4943abbc68cce7888db1ce",
		"fceaa8c636644207a08030b5584318dc",
		"07da6bfe3f2a4283b7b898c42d7a1ee0",
		"999b2ab08432422799bc521ad4403e44",
		"3e88e8d663214292a86d45f6e2fb2647",
		"a8301cd8850c445b996a6c20143c4fa7",
		"5d77f7ebab0646f2b5da51e16a08ee29",
		"0a54d5ce62ea4fdb8a1461b3b9cb9680",
		"de5f4da598bb4790a2d73661a9d94e46",
		"b6a2d27384104786bfad63ca1c79d933",
		"d8adf6b9ba434bb882ff61717336dee0",
		"e293084c1c034114b567e8f1963e1bf9",
		"a708d90f0d8f4fad8d4c7fba5ceba83f",
		"5709b89cb7ac461e81000745f960efd1",
		"6762957bedd548ceab251605f7ee29d5",
		"996bff7c6c11437dbed698ee70d151fb",
		"7c8aac90f83943a78ed643265fdc8464",
		"80ce4d9dd1d74f4096197778edca59e3",
		"10c406a2b3454c8f96d77248a40d2423",
		"4134e2b6cd4b4a38af8362b610bc1c59",
		"1367ff80989048f5b01c7edb40bd1192",
		"ba760871e91b47b0b2657b30a42ef0da",
		"110c57d1515544d38d03e43e512e93c1",
		"88c8fa1066b448a9920bb2b0d1046b87",
		"a1526f8e48f14a8db680908d8c943e8c",
		"1ce4a5297d1c49909057fd1f27a27b09",
		"94820992ef4744eb97ea3b64a35309c5",
		"2891ded5b3d649518c1f31f518dbb700",
		"fd39a1e1265141679526a2b8a79cf258",
		"59e1c829c0f64f5fa092c0d6c6bae0a4",
		"76e8ad29fd414c9c99473e60004769ab",
		"33cd4ec853bc4fe9bc998f68b76678f5",
		"60655f1e2e454bcc965480662e6a4c3f",
		"2087be9df91a47ddbdd13dfa430da3e1",
		"76d14d92192f4e988836e383f7959396",
		"74cf0a8cdab747b1906c584e2bf0f351",
		"7611e685d48344d7a10e759d6e52db2b",
		"b787f76ecbed42c0a508cee3cbbe3a17",
		"96a6a93c3c804c0d90918178255120a9",
		"f01a53df0efb49e1b762e23fde03844a",
		"f6099bbd5b444785a31dc09be3f902db",
		"4742c05239b54e5d828741e45ca67e61",
		"973a5ef66b864ab3b705f3683e9cacb1",
		"a5b793c32ddc4ef3a96d8c526b7a8066",
		"45f2862f6852414bbc8eea1c25b30113",
		"6fe3a060e98c45fa928b916d25ed7527",
		"4bf4762b96e64f4f8da0dd2b2793bb61",
		"c421bcfcfca24046a33825a07bee4db1",
		"441721d3a69e46ec928709e8a5580494",
		"96d7b58d0f564cd59f07155dcd0b83ce",
		"cbe8175f7db24eb1a973c268fd026fe8",
		"125a91d8e8a4485198de71e6660723aa",
		"e3d1f3e8b21a427baeea0450f4cf679a",
		"cddf7d68e40b4532aad9f1228cb630ae",
		"9974483fe2d743f680a35a6867c5df19",
		"b4db96e1a1764fd797f0fe6283bceca7",
		"b3d79489d55d42b795123f84f72c8194",
		"d11839000b804312a71c645168231ee8",
		"184c88e67d864f939490d9e9a3f5ea07",
		"34c3f7ea496b440ca8d4341d36795210",
		"efe69837ebea49c293327adba99848b6",
		"90156730fa3141b7b8b7c52df927af5e",
		"5ed9656501a24f6da5a25da54bad6da0",
		"c3aad171d52847409c19501808482e54",
		"a39b36161b1e44e9acf37ac87179f3db",
		"eabc481029b6467c9040f51309e8434f",
		"2aac43bdebd344f09825ced1222d5e57",
		"4eb878199a6d4040a3a5ed629ca1e399",
		"4fac98e7ff1d4bee8a94b7744147e180",
		"be908846e6334ad7a1cdc49d0071e5e4",
		"97fd8bbc7f2342299e2c539233e04cad",
		"6cf6f21c1b1b47afbf82280d3d71eef6",
		"823546c5ba0f411cbab9c45366c2a70c",
		"4375a98921b94e27a8ba74d2d45decc1",
		"0307ccbb543f4e7a8c985f150236550a",
		"daaa50e6f81a4a019f8146ec799c8076",
		"a28eb0125cb145cf9d56da765a40f0b4",
		"33eb59e7b85d4e16bb63aa34c9b58ffb",
		"965e99cfb8d94599896a68504c084947",
		"b1827541ed2843f9a87c88d1930c23ad",
		"2b780995194e42568c62d0fb280b7527",
		"94abfffd71014fc7be24c1e90029bbfd",
		"a19a56f1e4a74a64a92774eac601c74c",
		"e464fe8c814a44c5b5f0e7f9be594ffd",
		"b06906eff1124f0c969b6b46073eaafd",
		"862a6831a5414c6b82c79e1f98ea7cf1",
		"7e6a52eedbfb4319ae5ea55613030bcb",
	}
	buyTokens = []string{
		"4d2eac8cf1384ec4b699856e030d036c", //yty
		"57099c9dfe074484829ce872aa67e613", //pz
		"29720b3f8529452fbf2831f738d2a9ec", //zqq
		"e87709b4dee94ae794109deec9058f5d", //sq
		"b05ae67513f64651a003627e3280ffc6", //ytf
		//"47aa590705994433975afbe84437f451", //myr
	}
)

const (
	b     = 1   //1是分解 2是置换
	actId = 743 //活动id
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
												rwMut.Lock() // 加写锁
												//if len(orderInfo[k].Data) > 1 {
												//	v.Data = orderInfo[k].Data[j+1:]
												//	orderInfo[k] = v
												//}
												//rwMut.Unlock() // 解写锁
											}
										} else {
											if Replace(actId, item.OrderID, k) {
												fmt.Println(j)
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
					time.Sleep(time.Millisecond * 20)
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
									if item.Type == "prop" {
										if ReplaceProp(actId, item.PropUserUuid, k) {
											rwMut.Lock() // 加写锁
											if len(orderInfo[k].Data) > 1 {
												v.Data = orderInfo[k].Data[j+1:]
												orderInfo[k] = v
											}
											rwMut.Unlock() // 解写锁
										}
									} else {
										if Replace(actId, item.OrderID, k) {
											rwMut.Lock() // 加写锁
											if len(orderInfo[k].Data) > 1 {
												v.Data = orderInfo[k].Data[j+1:]
												orderInfo[k] = v
											}
											rwMut.Unlock() // 解写锁
										}
									}
								}
							}
						}
					}()
					time.Sleep(time.Millisecond * 20)
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
	if diffTime < 100 && diffTime > 0 {
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
