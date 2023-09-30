package api

var (
	TmpTokens = []string{
		//老周
		"b33d6005d6bd4cb6bcfd459a0a14e8f9",
		"266687b3f5dd46bc9b62a94be443a1bc",
		"54edb4cf4ae8489fa33a996b8d8d70f6",
		"e98009dd52b64d409b880435225e4ec7",
		"b69d0cbc6d8a4d1f9e1f54c3fc0cf5e2",
		"0c104d9e9acb4914af9594e52d2937f6",
		"764c670817f54ff4ae66f42fe730f178",
		"c4e2012afb1d4d9298bc608ef907e8f1",
		"e80572f6caef4888a31fdcb5ebd180e0",
		"4da5f9663dfd46eaa2d7014cb12cd2b0",
		"50ac8fcc1e64413a8df87415ba2ff354",
		"b4722a5b153e4995a56a28f0435c4ad4",
		"49168be055ca4e1facaea68adcb6b3d4",
		"426f112405264d80bf8b5ed914918e53",
		"6f64be79f80045548af7e7d35ae8f270",
		"de8f18bc08044bff8dc15e75067bd02d",
		"bb113b46885e44e89a1c3a6e634f78be",
		"029c15b20d4b4e2da2fb30d8a134e256",
		"9f4a75657c254edeae616941a63fe202",
		"43e3f5d57ef04a87af032b7900b0be05",
		"ba5770eb354f41b9bb172c8736150d77",
		"4c6901c2f9344089b3e4de747ed912c5",
		"506037d343b944438d058ba043045a8a",
		"deb67356c1204353983e410af519c72a",
		"d00acd2b51b34ed094c875bee77bfaef",
		"535b368f983e4e8ea06dd7899eb5e16c",
		"32347dde469149209d833bcc811ee2f6",
		"d101f4af6602449d96ffdbec4dd7606f",
		"bbe347fb36f04f52804d8732443543b9",
		"80da82b79e1944e8a6eb1871cdb2c2bd",
		"629f4c8e865e44ab955f18ef90c30040",
		"f442cff75d6d4b4cbffcb519a0d8287a",
		"cee3aa5c10744534b8a9403f84b7d640",
		"0ec2d7abf0d84553b672a9d0b9254119",
		"939a812e0c744f10b53be95babc6f482",
		"d73e76ff686e4fa0aa842d715015a694",
		"d3ebe5eee0c54ea1b19ff05f86397543",
		"a5be1036ab7442ae8f5ef90a0e73fe1e",
		"e80b77f63c6e4ca7b78355c4164668b6",
		"2a26dcc9204a42a28461b50c40bb7e13",
		"f2a78c70e5164f1b9fc3c7fc04713ea5",
		"9987cf3ec44c42619bd58dc410d14754",
		"99059e268b4b48aba5d7f158b2570f15",
		"7d3729ac3d1a4c4082490116a1c14408",
		"7e9d4d1e1b3d43abb982ffdfdca0040f",
		"aa29ba752dc54016bd0331a7633cc06a",
		"2116224d68ba4e058ff750709019d3a4",
		"dbfe344053b1417184194416bec851d4",
		"982ee2d3d18f4e5b9e1ea773ac444c84",
		"f4d7dd051fe24cf98bc29e294d6268bd",
		"b6c51269e6e742099f21d39755d29483",
		"203c788e19014efc9215e67ea99040b5",
		"c672c0f0033b4248876a38466f4a4c0d",
		"ce58c9310069479f96f69093f99b3193",
		"a2a5e9af46d7457ba76fde5f0e21e2e4",
		"d631ccbe13664b5e9fb72b6705e5efb4",
		"e6cf5069d53a4528b6a7f36e9867bb9e",
		"2617f1d456054b9cb573f82330c6c146",
		"dca173b5cae6444cad5c8293a68b08f7",
		"9ef29250fba747109e4d46723087c5ec",
		"19e87ca0ad7e4d5dab6e67265ece07bb",
		"ced5fdc13cac4138b8389afb94705072",
		"c963e8d508a34fc6b302ed77258678d8",
		"f9386cd28615405687db62ea79ca5bfe",
		"b7d6a84982d442eb85f26ff12407abbb",
		"4803e503a2264ce380b2ae7388dbdedf",
		"9b6f97f741d84b7d955ad7c76efc659b",
		"82392f35d98f4cc59c22c857da97e02b",
		"002495f4ba5e4130885c174e99a20b79",
		"da49c9246b4944d59cb733df7b3e32c0",
		"f04d74703fac4a7eb13c7230b8522d49",
		"4f126fe91e9c41ee9889490e1ac780b7",
		"90faba3daba94e399a96e4dc55e769ba",
		"d5839e3d10104381b5f95107fcf3c066",
		"6444270f9d774f8a8c63b1cd0636bdd8",
	}
	Urls = []string{
		"/aiera/ai_match_trading/nft_second/sell_product/list",       //寄售列表
		"/aiera/ai_match_trading/nft_second/sell_order/pay",          //下单
		"/aiera/v2/hotdog/order/prepay",                              //快付通预支付
		"/aiera/v2/hotdog/payment/kft/confirm",                       //快付通支付
		"/aiera/ai_match_trading/nft_second/sell_order/wallet_order", //零钱支付
		"/aiera/v2/hotdog/search/product/search",                     //搜索
	}
	//BuyToken = []string{}
)

const (
	TimeSpace   = 100     //间隔时间
	BuyNum      = 1       //购买数量
	ProductName = "武神赵子龙" //商品名称
	//BuyToken    = "65d7760dfcb14c449ec27d85ba93526d" //购买token
	BuyToken = "29720b3f8529452fbf2831f738d2a9ec" //购买token
	//BuyToken = "0db20208ef454d0aae633319b4863dae" //购买token ytf
	//BuyToken = "b86a373641414866912d2cb93c71f6c7" //购买token pz
	//BuyToken = "fcb5e5ab940a4db5bf7641b3ad6fa16c" //购买token yp
	//BuyToken = "406fed690c9a43f1a151cf2783a24561" //购买token zyw

	PageSize = 2
	PayType  = 2 //1零钱2快付通
	//Pwd      = "DVqBnIG8tFOmfbFp+tIXisluxkZDahm5Gk6MVvg4tY9td7tfjTvu5JiCDBmW39mUhgjY0z6zzlfj6Jc0/YDyaGLLB8n/wRXHoPRv6qlOyMleQw1iU5Y10MfF0jYylh2EJtiVd8VQWwOWgAuYmCIYUNqoy4IhjYxMs9Bj82l/rts="
	Pwd = "kSmIyvNXdGWlnTJBfmhtkTmvJc/sB8Bu78UrgHpj4+I3DsNaUpLRQccpEaBpKpG9+DeJOsgCwrK9iL1JG99GjUxaL+loiCDcm/UM1EZDPmveXq1XQeyLhMjbyPEGGUnJUp0fwZUAddRpQ9Zanbq6gk4/lEftFxFWz4wIHsfgJQI="
	//Pwd = "oYM0g+TEQ7ToqmcDXpBKbEspk3H/buLHmBNu3Y5mqFuVPx57gjznAI6HYzCv+evM1DbKjdWHmHirGX0KhNkuk+Wjp6gw+0Jz2AeRLz8u5ChQHv+SHW9Ff1CaNUOyusBtJcg2JDDA17KCaJ0vCFYjVdQXJCX+BWWUwBP7zjHrUNI="
	//Pwd = "JLI38dg4HN9dAnw89eChytrWE+PmWQfxAf8u2tgIlAmZ2u/WCxYiB0p1MoeMH/qosVwXzY53Ck+zfK58Ox8Sl+MY+0kGstkApiLema+rAbPA0inMU6jlFGTtpUOKHaw0Wts6AjMTSG6aXyReh/3LuZFQfcGLyczTlF1egqZDsgk="
	//Pwd = "QcF6cbAHoMiEfH8okZX5Jug5wMI6gByX24Gj2m5rQBGpl/bqT8VLaIY982QtyIMkdyTz2SeFTJKPatFm8GFWsWFBkU4xc0lyEJiG9qkWjoJAdQ+R9gP8g0tfRy9cWrRtaKuEMhMd7dUQtIjoxo1v7xuRMyHCLC3NLywZSu8x7rg="
)
