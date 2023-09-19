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
	id        uint64
	orderInfo = make(map[string]ResponseData)
	tokens    = []string{
		"cac29068ad1d45db88eb410c0ecdbafe",
		"4ece3f0db50148ecbb59923b34982f4f",
		"1cdb1557859148b187513879e93acfa4",
		"7f814785a5b5446fbb65f90564115d9b",
		"2a62e05347454fdf9d82c0b73f5eb5ea",
		"69a153491cc549c8a679aedf239e6d87",
		"8dfcb52533c949df9595613f25e82f85",
		"ccafbbd5a5f14485b7b3864e3a1acfb1",
		"265b6448e0824170a078dd877496db58",
		"f66e6f14ed5747d5b2895aafc34b85e1",
		"c06c766d13ac4ec4b659adba9997daa7",
		"2cf24880dce84da28deb5d48ca631103",
		"2398bbacf9ee439680899b28891d2f5c",
		"269defb055694470b8c9d2e5ddf24302",
		"10a7c1121fc748d2b651b53955dd3716",
		"78cbcc3d4cd040618ad3742e4f388f70",
		"ff47ddd17e7d45bbad521a4bddf93669",
		"77fc1d99f36f443b9306f322760e42f2",
		"7cf891e406884ee6978d97d3cc2654d1",
		"918055f31a114417b0dc0a4d7605e94f",
		"f84e0824a37548bab559c36eacce2d1a",
		"0bbb0c2ea84847dd90657f1318b44f9d",
		"14eb87926a4543ea961646129da96ffd",
		"f20fab7739e54e5c9a0e16566def12a0",
		"f5159dc7230841228f449ee7c419c3f6",
		"a35c893822634079acc15a0bf76aae15",
		"9a0bb077a54b4afe9dbf2c6df660f0ea",
		"0dc668d4616b49fbb007d1b7d26ab628",
		"c44951f74ddd478f8cdf7057e0cbe1b8",
		"f3d2c9901cbb4c1597183395883fdc13",
		"baf55c06c23f4488a6c9eac5282487de",
		"d494ee0ab1aa48b0b4680b927384d92a",
		"802c95c2d2384497b592f9113f4fc954",
		"bff60ec7b1d84674acb54d2bf9e92616",
		"f755a66b57124b8587f7d7d7c37bc780",
		"af2fea66e9f14a3ab5da7c0062b255be",
		"d140e1ece17c4bd1b3936004dca4327b",
		"605ff5b45fb34ebfbe07234f6f4e5d40",
		"9e698e471529405e8924fd92f9533c9a",
		"392962c592374a5a8811223039bd918a",
		"04ebb5abec3144dd8da99be8ab8b41ca",
		"1d5dba394ba34b2fb1a67e561b90ef1f",
		"12416fe9ec524ec982ed68d838861280",
		"aedff138a86c4528882edd662523531e",
		"2f8960f930cb4f6ea68364c77e819c29",
		"abf9982b1fd34781a3df8ec7f150855c",
		"b7a0ae88626443beb57f5742b94dfe93",
		"580eada36921424886297ca14ddfb459",
		"4c01f5b22b7f4b8fba89de072cac6440",
		"d2871ee574414d7d9f129d60c79700ff",
		"4c5f2690aece4e988c278585790ccd34",
		"6b4f209ccb64468489780ebc04636070",
		"afac65f80e7b4321a29166b39547a921",
		"a755d965554540a3b8e2a012ea74608d",
		"4fa3d8fc48b242adaf3d6f660af79d84",
		"41033bb4ea3a4f13a81a402e03aef049",
		"34ded9f3ef8a450c93b442ce66a80aa8",
		"cc6c52d9065e42df9a9805b25ce9ec3d",
		"8299101598864e30869363c2fb1e9723",
		"ec12f5b7dbb84408b3db5a9b958db206",
		"1c3ec550483e4ec2b66ea9566ffb7bf2",
		"7d18d75cc4014f7a93dfc47fc9626013",
		"62dd0ce5b2f34a5cb9808413655ec9df",
		"839f3b9dd2e64473aa41023386107513",
		"568c16b3316c41cdacc23b44ce0959cb",
		"996329e617e24d1e9e059ecb18c3a85c",
		"a62b902516a74cd298c78fe96dc8acbf",
		"a651b1126a76425da286eca88edacfd1",
		"2998dc58699147dc90bc1a8636f7a54a",
		"d6f53f01b05747aaba75efb3b955617f",
		"17e93ba4255843cf826a83e10a5bf1b0",
		"03ee887154484ea2b6dcbcf203255711",
		"f7c8c668ad0c4f729654ab95294237a2",
		"0bb908fd6a334d51948c202900549431",
		"480e716107554d17833fae7eb4eeb817",
		"8985b42f2d0a4c4b81168a13dc8b494d",
		"2c7e12cf54da4489bd65d2aab8a0c0eb",
		"b0daa83185ae4a25b7064f5bfd5ee4b7",
		"50ed823fcf1a4bfb99db87ae15f73c27",
		"78d8f0c53eb340ef9259d3fd4bd84df6",
		"25661df20ddf471594e095ad50e7aa72",
		"866914e621144b068583526d03ddcb7e",
		"a5cfc0c94b914c7b90e03b70e8dfe742",
		"78b45c035a124ab2b6edbd3e2b8fee86",
		"56f21549681d42df82b141be790c1b7b",
		"493922ea35a1499eacdb1ee45784e0c4",
		"ee41704a5e8d4096a97194db815e4d05",
		"9b7bb4f02b1e4bf49b8d7bdc7d3f85b4",
		"e2f2d1cc444748bc9c34f080601a5eae",
		"986839edd8dd417f95a03974551bba1b",
		"445eac4249c64056bafa4b5f0bcfa5f6",
		"17cfb3caa55342d1a2cb6a158517ffa5",
		"3d40bb5027754eb78de8b2d869e556b1",
		"a7ad22f04332479180f549e0eda1153d",
		"cd0022c2506d4a0491dd1d67e32d86a3",
		"3b4028f9ced0447cb81cfb156979bdd1",
		"e96e08122cd8477387fd872a8d7ba885",
		"47cdaa48f5884c6498361ad192e930cd",
		"fc49f8924c5a4729af0eb8dcd84d9d29",
		"2380412451a942fe92376be32dc849e5",
		"a07500177952498aa8d313af14f659a1",
		"84b3f40734d64a8ebb5a887e1c89da37",
		"28405c13941b4267b276a621c9904e03",
		"47156be9546f4e1e84943716a780ce1d",
		"7665d4073c4349b19eff1a241743c6c7",
		"22c981070bf243819b36e27636a5f440",
		"c7c496a7d1c448169f0e1d0dfa0e6f0e",
		"cc33c2108b524f5b8d209059a0dbca4f",
		"383489fbbec142dfa2b6a4f53efc8665",
		"ab998e8fc0be40baaed9edb73c749408",
		"27efdb787a3042698a39749afb9b6497",
		"2f7f0e3f940d4edc9344783b605b839f",
		"bfe3978f2dc44a93ab8884eb9833e861",
		"5b1ef15ef2f3403cbe4db069d213eb97",
		"fadd11808d6a42ffab215473af79f4c2",
		"7c4127884bf04ff2ae39b1deb154f621",
		"e8b24000694e4755abc017c00a5e7152",
		"04493d15967743e9912f89226c1f0179",
		"4be2d85c99104460bb5c4766b2e0c263",
		"22774aeb52b240e39d470d87f02ef139",
		"07f18fb678c84d7e877b1eacd45ef70d",
		"65d57b11f4974f3e9c98808c2550115e",
		"feb47f3ca1d54725a21cfceea02a95e0",
		"11bc728276df4e41a3aa780dd7738122",
		"d8d3d014fd7e4b4ebdc0677e4a99b4e1",
		"ce2340b125184e108acb74fa6b4d5b05",
		"f593a3b33c144ae8ad93912e058c2b5a",
		"7541e02d89334f499c81850eab393383",
		"b5979973b7a24158895f518f80cca8bd",
		"d283dcd98d2943468f6fc2d572a0bfd6",
		"6a481ef1244c4b50a6501974b776c497",
		"16a114fe07a44982a3d9c2a2f0fb43e6",
		"7b9dd98680224c3997608fe106a7e409",
		"9fcdcb3ca3e44b99b90b5312530a35e0",
		"f0ffc1e9efa84bdfa3a0fdb556fbad56",
		"e039a9e521414fb99e0d6701372bea4c",
		"1432120246f343ca9690e89bf9f3e2ff",
		"3081474708b14a95ae312e28d5bf584f",
		"b034e968d32646b793b614b7afbce98a",
		"043c2c42729d4ac797b355d27f5d1a03",
		"ee6e9dad5bee42e9837634b9e30264fa",
		"21106a50cd544507b259062ca2a8c3a4",
		"f0063fe462d04ad690a8136e4d11e236",
		"df7b128b3b45494bbd12a452bf84f623",
		"939293c845ab46d598547ce1cff16c8e",
		"72634006ac8341ada26d5c1dd62ced9d",
		"492add2bc2994d67ad4ca21d082108ca",
		"798103ced0724357ba80b526ad75184d",
		"04f3fa80067a4273989f3edd41e58a41",
		"0c2b35845f2a4c6bbc2dcc6ec4f22370",
		"7ec7bfd30c234a67b8063c0223efe6f9",
		"976660617b644b129d47fdb124e8c501",
		"84efec1d4aa84729991ba0300f0e9ac9",

		//zqq
		"0cffa5bbb48b40bba599c40e7e407e0f",
		"3f6afe3f3aa243bea7cf44bb21065efd",
		"f8b5fd685e97489ab88541cf431d79b0",
		"a609083331cf45c995bae3a31fd97c29",
		"39d45a53f11a4658b9a483de9bd42107",
		"06a4ed6f6d2f41d9b8190e6b95f5094d",
		"a3245f6bc3af4296a5242e8f5af7b86c",
		"33efa8a7c4804900949879f2ab0bd366",
		"2cae0d8e5f78473b858a955f810de3e6",
		"57d855f196a74c448640ea904d10a468",
		"f6e8cc90f76b405989f3a7b33e8d7914",
		"3a0be1b165174a168468d447dc0b1a81",
		"04787bc165f343e093e955270fccecd3",
		"bb6cff47418b4b56979847f82c93af96",
		//"46037b1ee23e41aabd312f0aed422962",
		"3299208b97254154942b462028b5e5c8",
		"10f8fe9ee8a94d1a9b9df63439486822",
		"9c096c73d92d4c778b172d275670ff32",
		"66fa447413804759a5f9cceb0001818f",
		"aaefdb11f16e436ab56b04f1a6d7233c",
		"7064683e0e18496ea04ed9efdea1e8c5",
		"94f53ed72673403eab51298ec6b6c9df",
		"b2fac38fa81c46c58cdcf5accd559fe2",
		"8293552fda624031a0c1b53b4031da25",
		"bdc059ca4d964f7ca266b2169fbdf2fa",
		"96621f2894f2425d9600bc0ceb41bf86",
		"a1c09a95b6c9462e8f460127edc1fab8",
		"5c7d3c5b07414895b343f1023f0badd1",
		"90689dcf29c14ce6b6df3a91aabcdccb",
		"2a080d2d52a145d7a2f060ad1d64b8fa",
		"7dfe13158e664d5abe55d60d548af5e9",
		"282a66f657554d4db73acee617ee58d6",
		"4f68f9759e8a442d823c4afa0975e981",
		"9263f490e8f34a7ea46b37fd5528dfca",
		"27c91dc0a9ef418fbe82c19f716e3c61",
		"ddb60f64564b46e0a5ffa00aa4a8c354",
		"22e63ca03e954e36b71a39774c523f63",
		"0397e7064ca14836861b3d28267a7838",
		"c93d4218eb23401890c2fe287cf8da5c",
		"6c9c5107385c45ffaa4383ec8c50e478",
		"be796d8e035f472eb9f30fd314d451de",
		"2bade835c2f640ed9ba2a8c27ccf623a",
		"45cc8ef91066492493f772872ca7d2a7",
		"c9ae3e946d694108bef9f3e136dd5fff",
		"4731ecac914347ef8cdedb46ac8bae0c",
		"e87a27333ee7433d9f9b70b9437fefe6",
		"fb1d12ec5525480b96c3e330f64236e4",
		"2b3b662b0997442eb715bfd8654d82cd",
		"1f3890a3dd9043ccad3872e2836a13de",
		"f5bcba11a31142f5b183964c5c04d73c",
		"e79ff4956e7a47ac9f75560bbd8b273a",
		"586c95f1a4b94ad7ab9d8784229497be",
		"c94cae76ce5d40c2aa4bd066fc67fc18",
		"feed793beae443e78d8cc24eab44d478",
		"4ada4821d0d2440baf046eea4a64e224",
		"5a0a404c30314a02a1c341e57baa45aa",
		"fb9208778498427ab3c5c954f81f281e",
		"b89c8c7825744bf7b0604339d1cb6701",
		"516ceeee249a4225977013509fe64bf7",
		"4696a38b76184ef1b58ffbab68f343a8",
		"5ee5eb2f108643199ec1e8fc1411a4fb",
		"63503be8025047bc96a7e495fa1eedcd",
		"5a550e8cc8144a29bc79d6cc4973a87a",
		"3fa753a89b5f444a8942878b50e3ccdd",
		"6e336d6ac1e54c7ab4cffc9ee75b6073",
		"6e56f17c7339494b8e25f9cc42e48a57",
		"d058b8f8e81c4306b39c770e905245df",
		"a45e7f1516e647efb990c7caa8ad1150",
		"4ce7f3526cd148c1bf1bb41cf2820156",
		"3449fb9b21cc455d927508556d6ab657",
		"60d0d881a9de4942bac2f222d83c225e",
		"2a6e0714254b413db2b24ce97a0ce737",
		"774e770e53c143129ca6f7adf26dc116",
		"406c3317f01a4c8eb8d718e5729eb037",
		"cef71d498ff74604b8b607e5db6d632b",
		"2b504fab8daf45ca8b4634969e8e6c9d",
		"0d6cf9edfd1f4abfb1ae3b1a54077515",
		"fd06d19c68fb4faa8735445170afe5ec",
		"e32f1ad284c84c5797f730cf90d0833d",
		"0c1eb7bbb0124b28bd910767d2c6a8d9",
		"d093342459104f1389c552e4cc74d01d",
		"5fe29770e36d49c39e294ddc83db1b2a",
		"71eeeac88bd14fe69bf81246bebe9975",
		"f5ba47691122454c83b23602dbcecbe8",
		"6761a7daa34944c784e69134db352546",
		"a0da486f0c994426b5c32ed122241709",
		"2731ef83b1ce4986862932674850f76a",
		"448e3ee022884e639290748c99b10fb8",
		"08a2d1fcde2d461bb69e7a5d9ee6a3ac",
		"2aa9931d493442389a6d3f8f37dd0aab",
		"787dfcb0c1504e3cb5144ea024cb12ce",
		"a5167c8a4a104b479906c1670375d974",
		"063bce8b982c4461946593600a5ccb24",
		"e7a30d8290e944e9ae2cd555fd842d76",
		"e4dc18349f3444ddadc807fcb0ddf499",
		"52c1875f4b2b45e5a8f4688c262dfc87",
		"c3ff9c40e12c4bcc8f38a2404b52750b",
		"49720b01bce34afcac79e991edb99166",
		"c41250d0504b4150a7d870fdb3eaddad",
		"987c39f635394272a4ef6dd9653bee6c",
	}
	buyTokens = []string{
		//"da01634063c446659313a5a1e013f86c", //yty
		"7b50884dc56f4c779dc7693617a7cd8a", //zqq
		"b86a373641414866912d2cb93c71f6c7", //pz
	}
)

const (
	b     = 1   //1是分解 2是置换
	actId = 621 //活动id
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
												v.Data = orderInfo[k].Data[j+1:]
												orderInfo[k] = v
											}
										} else {
											if Replace(actId, item.OrderID, k) {
												v.Data = orderInfo[k].Data[j+1:]
												orderInfo[k] = v
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
								go func() {
									for j, item := range v.Data {
										if Replace(actId, item.OrderID, k) {
											newOrderInfo := make(map[string]ResponseData)
											copy(newOrderInfo[k].Data[:j], orderInfo[k].Data[:j])
											copy(newOrderInfo[k].Data[j:], orderInfo[k].Data[j+1:])
											orderInfo[k] = newOrderInfo[k]
										}
									}
								}()
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
