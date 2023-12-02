package api

const (
	Host       = "https://api.oneniceapp.com"
	BuyToken   = "rH_R-m5BVEs8r_pkgjy-PXTb0gQjyTrT"
	BuyDid     = "d5401cf612846e7cd15a2318039d67b8"
	product_id = 852926
	TimeSpace  = 6 //间隔时间
	BuyNum     = 30
)

var (
	TmpTokens = map[string]string{
		"rH_R-m5BVEs8r_pkgjy-PXTb0gQjyTrT": "d5401cf612846e7cd15a2318039d67b8",
		"jkHmrdtnGqzpIiHeg2XwPX7jxDhRKZMt": "d5401cf612846e7cd15a2318039d67b8",
		//"nuWOco5j4NzwJwe-fkDZPoJ8279jYnnC": "2b97c4dc689f2df3e26104f807438379", //zqq
	}
	Urls = []string{
		"https://api.oneniceapp.com/Sneakerpurchase/priceInfosV3", //寄售列表
		"https://api.oneniceapp.com/Sneakerpurchase/config",       //确认订单
		"https://api.oneniceapp.com/Sneakerpurchase/prepub",       //预支付
		"https://api.oneniceapp.com/Sneakerpurchase/pub",          //提交订单
	}
)
