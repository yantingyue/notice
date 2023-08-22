package api

type SellListResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Res []struct {
			SecondId uint64  `json:"second_id"`
			IsLock   uint32  `json:"is_lock"`
			IsOwner  uint32  `json:"is_owner"`
			Price    float64 `json:"price"`
		} `json:"res"`
	} `json:"data"`
}

type PrePayOrderResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNo string `json:"order_no"`
	} `json:"data"`
}

type CreateOrderResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderId uint64 `json:"order_id"`
	} `json:"data"`
}

type PayOrderResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
