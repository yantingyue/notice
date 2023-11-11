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
		//周穆浩
		"05c43169ca834ee6bb01ab091bdd2a56",
		"49ecb6f0b53a4c149d9644032f6bf73e",
		"b89ef1abcdbe4a8d859c0c00e0691354",
		"499c68a55e584a87bde0e95be4c526d2",
		"ec3882b5a8c94d638ce30339890a387c",
		"723d90e4cb704cfdb58cea2cba4e96f7",
		"2b29ddc6c3fd47c5a8973000055950c1",
		"7a55cd893c1c4eaaad98b423db315e1c",
		"7ff20057b2644d88b0bc59d97ad8e8e3",
		"7c5f873a51164a0997cfc18002a373e5",
		"c6e3ffddfac24aaaa4138dd14f42cb57",
		"2c825d86639841c1b64d6e4b27da9c05",
		"20908d87c481477ca648020d1184f487",
		"33c31f403b0947f0ae408c60fced07ef",
		"b0890ef257704b21a82b084a00e96dc9",
		"8e366f457bdf410ead9f6518e2020a09",
		"6d2c2dfdcabb4085a679cccf083758d5",
		"fd8908baefde4a4792aa2e54ffa19953",
		"ad3e3d05009248cca0b45eb6fd4fdaf1",
		"755aeaf140b14f3ebdd0d254c3772e2a",
		"17ba81b36fd041c6857cfc4fdcfb93ef",
		"3c8c3f2f964540a5851daf5c54435f48",
		"84f9b04a39594da19efd62a97f2170e8",
		"b8d2281fc732433a9096b7d6034c31d0",
		"81820f541af24f6abfb2bcea4b501788",
		"85610487a44f4c04b90d775079469a31",
		"f25b54f7180d48209172921fb4099d6c",
		"7ff7588b8db54d3990298abf58c28ef5",
		"5008a46f9a824668a3545f6a2a9a6dce",
		"bcfb7110219f492980d8497d535b323c",
		"af3014ca3b3945cfbe38883e877fae69",
		"43dcf10d5eb24f5f88a341260418c2ab",
		"46b7ef92aea34a279b47014d5219a7dc",
		"43afe5ba80a245078649777a1faf67d2",
		"4716da4f12444eaab134fa3cd40972cb",
		"779f59ea4d344785bc5e63c40a83178a",
		"e0b14a06f7c440d594b100df5a8df152",
		"eadfa862018c48a880859829925760e1",
		"4ab0c987b40b44d5a3ed77430fa3d884",
		"d3caa203f12243b3be92bcdebff8f9aa",
		"3bbf3fb392504692989f326ce27d1f48",
		"0746da069d514a429661b7b67359c65b",
		"b09623f40f864958a1e2fa36ecd4b0a2",
		"fa0f4a31ce054e94b261d3dad75b7430",
		"ad3d10545100473aadd563c1be5b783f",
		"a5510e6d144248ff8315631e54d2dd57",
		"549d8a5a52cf40cf92b99b54beca15c1",
		"442a8047777449c89e190fa9e57f648b",
		"524b104a24f34d0f85b64501f7c576e0",
		"0548864cdb3842ae915a03a29c3e68d0",
		"9846c1a00a504c6081a3f244f0b3688f",
		"59f72a7acf364f53970d853ca6365eb2",
		"a27ca91d945f40feba7ce5af4b8561a6",
		"9d32892587504725b9a75bc295d4ced5",
		"3adecd31fa6d4c0ca4f26a57e5974a82",
		"13b841c88d6840239995651a4c53c772",
		"bb5b738ac3f34008bee55a5bf988d63a",
		"323ad4d027f14dfa94836e1b2832f6b3",
		"b35fe230c28241918250c36c662a176e",
		"63f093d68daf4a0b9ebf8f2b0305b833",
		"5f06ae34c4a84dc982380776d12555ae",
		"5d8279d643c546fd98061779f09f4319",
		"d739fba034f14c8bb813692826d9f858",
		"3aab2e8a0673431a8611b0bb31447e32",
		"0e7f31a6847d4c61ac439021995e07d2",
		"d5a47c3e6d674650963675008be06401",
		"ce27153dd9a243af8f05c1a76406cc83",
		"e5d6aa3658834c80ac6ad7586dfd1358",
		"56bb6b727a1c4d088be9deacd55dd8e8",
		"d54bf0ea3c9c4d3ca4a3a9c595d43ae5",
		"3211d18187814fb0a2507783ce364b75",
		"a64ad20f62b84dc6adf072a9e95fbc9a",
		"f83ae875d1704defa3eccee21286124f",
		"b5042118fa5b4812bcf0242924573c0a",
		"c604528e045c4e7daaab417ab96cc887",
		"cf2f5ad9b1dd4e8a98f19f5117767feb",
		"e0a6492ac1314bcbbdf90cbe61f656cf",
		"6aec2e3323f64462887a5876b0afc993",
		"e8ec5bf8d1234e349ae517f2199a5e48",
		"6e379fd31b184290a5da901b388f9851",
		"c570074fe36f4e738424ca64294f7fc8",
		"aa02d9c6bb4e4d69ab325d09df6eb23d",
		"01f06678109f480bbe7ecc43c578c20d",
		"03764f27dfe74b8ba3f14c0946a4308c",
		"10ddc0584e084bcca1e3f759de79e21c",
		"a80036f3f1804cd3a3c5b2b92d167821",
		"5a19fc9a003443e0bfc96870669d7b05",
		"fe12bf9b866d4f439db3075b75086d33",
		"a07f27c5bdab4511a8907fc5da7f7045",
		"ab79e33354fd45c585a45c12d2218019",
		"c7ae93f3309948c8896cf6a57c2bce8b",
		"79769bb094a24b318152472ce4d166f3",
		"27c534e1b4b54a0097300f2ca0a42795",
		"99ea966587144b98b800ec5eb64bf3f6",
		"7834e177776b49ae9199ae7f1fe8757d",
		"510cf30b7e004a3bbab833489d41ebdd",
		"2c0ceb178caf4ca3a940ad02fd555356",
		"c6616d4cd79942d395711fbe10a46d65",
		"fef636a56a274ae5947cfef8fcf1347f",
		"e6be63042b0243f29de84d3d063d1684",
		"3fa43f186cb3480e82da2329d5dd60dc",
		"83594c7b20eb43f7a90dd7cf0c8477e0",
		"3ec74e68cafc46e98d5eaaf06115ea36",
		"ccdbafb4af364ceb9af06ea293a3a185",
		"517759c683b64118b45263004eae2df2",
		"5aa63c53f8454f81b775f3e6cfa989e9",
		"a77b2920b52241ff900384f350d9ee6e",
		"9b637ea954ef470d89f5f57c8c360033",
		"1a2c29c77ce4434c9fb8e961fcfb24a5",
		"1e69342f4d3143238375dad56be34509",
		"90960f092bd846c5902103c488f2f864",
		"0d60962b5cb54e98b93227adee6f9292",
		"61d2f911312a442baa99fc9eb195aa36",
		"096dd5e5a7e94ab282d99bb35ba787ff",
		"98047aae295649eb9b3710beea0c568c",
		"d58a7c43c765425abc12dfac1150eb39",
		"206fc2d7556041128d381e5ed3c2010b",
		"ea05240cb2224cf3ada106d84e89d348",
		"5ce2d31579b549d99aec19f7af5630d5",
		"cf8f8c477abb4de88a58a7b1752a1b2d",
		"aea7cf7640ab4e12ae5ec4a2cee2c2df",
		"cb7ac3d8cce34e2fb5d1f6bffb5192f1",
		"cda4f62dce2b4733a28939804792b17a",
		"ffad5e489d504d88a686facf1793553e",
		"140d80d9cbf3437199d50d451e35cdc3",
		"dcffbd35b66543fd87516bb625d6a12d",
		"f31793f138c74d519f8cd54bb56ea6d1",
		"f598864a64ea46ca83e361de7a96776f",
		"9e42e6041f2440cc82fa710ab1635a24",
		"60ffba23c49c4a8f8fcc610cabfae847",
		"551e5b6c3b4246aeb058df6ebe6dbcaa",
		"55562cce9bf34596b6e92f61206a31e8",
		"a56d157241e44bb2927ea5c4f5a26779",
		"c087d91cacfa4271a9e011f1ff22e8ae",
		"463e7e83612345c2952aca5599078b0e",
		"ed3cfa77e02a421abc60722c3e206d1c",
		"a3507c9bdfde4e9daa52040d3c39c8c3",
		"aa55220b18b04519931effac3d5f66e8",
		"b9924834add54526b5a9c8870fa9a9de",
		"c88f28b9cfaa41079556d3bfb29ae046",
		"792def9410fe42288bb2b5fd73f02f38",
		"d0993416d6d24336a84a6a1ae414caa0",
		"ae73d2e3e8b4433e81cc1f61522ab00c",
		"1315b567a16e444783b5514c9e69ee3b",
		"6d740ed1d1f14fe681f599f77423aec5",
		"bbd8d6f3cdf14fe6b2702f22cee00452",
		"048b79a5454b4212a039dcf1361b197d",
		"1b77a9a75c51444a8bacab09bad13dbf",
		"6475b58afb524460a1f8fef2ff69607d",
		"7999d32e59cc403cbb1a3eb4592ed3fd",
		"777acfdaa7ab4709b837b7bcde292516",
		"8f5853b68e3d45389c9739b7b14f18a6",
		"e66da6aef84d4b48aa1b92f3f81275cc",
		"d3ddfd89dd3f43e382109d09f7b879e1",
		"d3e1d15968ff4f97a193e34b3bcbdeb7",
		"4c0430a4269e47a8a139371558bdda92",
		"80606114d6574898951d7dffe134a5a4",
		"b073f5735d8844cba3fc5b89a14429f9",
		"8b86eeaa6ffa46578286dea66d462927",
		"6c64461cc55144779e1d2ab391d052be",
		"8e482f2e485345cc802b8b76271c57a6",
		"9d3f6571f2c5456f955af09a29087f25",
		"d5d3ef0f7c6543119adf58b735cf5d40",
		"e0b670a34c114dc6a07b87d6ac30677d",
		"3c1b8ac9adce4e4aa653c190c9ddb8d7",
		"bce8df0624294cfdaea6ff8c8140c94a",
		"01166f520b83483a8e59f3b774a6c975",
		"e244e849b4ae4bd39c6a0f802a349026",
		"322784d0160b407c960373264b80fcb2",
		"854571b44a1942fdb5e87d30e77f3d44",
		"1006a8d8870844beaa08cc284314b0a2",
		"fc24bd5ab20d4c5e98dba7fbac519507",
		"b6189627d7384650b35e0e7361f91905",
		"ae3a2339edf74f9d99dae18664dd10d9",
		"a9b1cf4a08184951b045328d8fe13ca8",
		"e15ed389e2ef4025962ed4209af906c1",
		"9233cb9c94044235866f5c99cd1cdc01",
		"3f08f1b0811046ccbd18d6b2cd6f3c2a",
		"f869c00a94814f5d9b58189c0e417dca",
		"9999a2041b094002adedd394e78dc36a",
		"089b6b4649524d72acfa91aa049e962b",
		"46e63b7a31bb4ca9826bca6048ca6613",
		"7d18b9863483416ead35cbd5648eec3d",
		"89544ad30cd64964a0476415d761d66e",
		"69e1b962a4b24af389b6f269243b033f",
		"012d79bd4aae41cd8de3be7f2a64680e",
		"caa89080c83441d892fe3037ffbbb2fc",
		"5b19e5d2aa3048e8b581588d34037fbf",
		"b0ba85d991f7481da2620b2b1e09f79d",
		"e3dfa8581367415cafbbc6e8fbe0d34e",
		"be494c0cd33b4585b1baf17f23857a04",
		"df37decb470d4f128b1c08fb56a7169a",
		"715d2938d2ee4112a1cf394d050ebc96",
		"d0f41d3dc6ad4ebea4168e6e9693b274",
		"eb63164cd45545ada93c6578d3441a8b",
		"c94d2900b25946488aa23096344fc6ba",
		"45ed83309f6748d88022ff271e10ecbe",
		"893d81103c38408a9a68e42055e88091",
		"4c8861a75b40427a9df3094d19b19b43",
		"1f46281b253c4df09c2cfd994bc1a619",
	}
	buyTokens = []string{
		"a99652e09e01473c95acf74426f60b86", //周穆浩-大号
		//"c7ac206276cc4510a43ce9dfc341a951", //周穆浩-小号
		"f8e6351734794a5395260118a7f9c5da", //颜庭跃
		"761c8046e89e45a6b6ed380508c3c8b8", //杰伦小号
		"7f0000f7b4dd4987a16f4acfc9449e66", //颜庭靖
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
