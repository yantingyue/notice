package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//api.Begin()
	//api.Test()
	//api.NiceSign()

	min := 500  // 最小值
	max := 1000 // 最大值
	// 设置种子为当前时间
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNum := rand.Intn((max - min + 1)) + min
	fmt.Println("随机数:", randomNum)
}
