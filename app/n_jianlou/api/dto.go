package api

type JsList struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TabList []struct {
			List []struct {
				SizeId uint64 `json:"size_id"`
				Stock  uint32 `json:"stock"`
				Price  string `json:"price"`
			} `json:"list"`
		} `json:"tab_list"`
	} `json:"data"`
}

type SizeDetail struct {
	SizeId uint64 `json:"size_id"`
	Stock  uint32 `json:"stock"`
	Price  string `json:"price"`
}

type ConfResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UniqueToken string `json:"unique_token"`
		StockInfo   struct {
			Id      uint64 `json:"id"`
			SizeId  uint64 `json:"size_id"`
			StockId uint64 `json:"stock_id"`
			Price   string `json:"price"`
		} `json:"stock_info"`
	} `json:"data"`
}

type PrepubResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
