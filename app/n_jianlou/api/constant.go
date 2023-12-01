package api

const (
	Host       = "https://api.oneniceapp.com"
	BuyToken   = "WXxr7urByMoVbnsdIsToPfTNwOl2W7k6"
	BuyDid     = "d5401cf612846e7cd15a2318039d67b8"
	product_id = 852926
	TimeSpace  = 6 //间隔时间
	BuyNum     = 30
)

var (
	TmpTokens = map[string]string{
		"WXxr7urByMoVbnsdIsToPfTNwOl2W7k6": "d5401cf612846e7cd15a2318039d67b8",
		"9i27L8kbQi-ui2eubJJiP4gAj_MMyy-U": "d5401cf612846e7cd15a2318039d67b8",
	}
	Urls = []string{
		"https://api.oneniceapp.com/Sneakerpurchase/priceInfosV3", //寄售列表
		"https://api.oneniceapp.com/Sneakerpurchase/config",       //确认订单
		"https://api.oneniceapp.com/Sneakerpurchase/prepub",       //预支付
		"https://api.oneniceapp.com/Sneakerpurchase/pub",          //提交订单
	}
)
