package main

import (
	"notice/app/n_jianlou/api"
	"time"
)

func main() {
	//api.B()
	for {
		api.ReqList()
		time.Sleep(time.Millisecond * 5000)
	}

	//i := 0
	//for {
	//	api.RequeatTest()
	//	i++
	//	if i == 30 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
}
