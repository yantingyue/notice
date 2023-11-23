package main

import (
	"github.com/henrylee2cn/goutil/calendar/cron"
	"notice/app/job/feishu_notice"
	"notice/internal/cli"
)

func init() {
	cli.InitRedisClient()
}

func main() {
	cronjob()
	select {}
}

func cronjob() {

	c := cron.New()

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("fuhao", feishu_notice.UserId2)
		//feishu_notice.CandyNotice("fuhao", feishu_notice.UserId2)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145", feishu_notice.UserId3)
		//feishu_notice.CandyNotice("145", feishu_notice.UserId3)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("miyuren", feishu_notice.UserId4)
		//feishu_notice.CandyNotice("miyuren", feishu_notice.UserId4)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chaijunmao", feishu_notice.UserId5)
		//feishu_notice.CandyNotice("chaijunmao", feishu_notice.UserId5)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("peizhi", feishu_notice.UserId6)
		//feishu_notice.CandyNotice("peizhi", feishu_notice.UserId6)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145胡莹", feishu_notice.UserId7)
		//feishu_notice.CandyNotice("145胡莹", feishu_notice.UserId7)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("门童", feishu_notice.UserId8)
		//feishu_notice.CandyNotice("门童", feishu_notice.UserId8)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("富豪小号", feishu_notice.UserId9)
		//feishu_notice.CandyNotice("富豪小号", feishu_notice.UserId9)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("xiyijueyuanti", feishu_notice.UserId10)
		//feishu_notice.CandyNotice("xiyijueyuanti", feishu_notice.UserId10)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("fadian", feishu_notice.UserId11)
		//feishu_notice.CandyNotice("fadian", feishu_notice.UserId11)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145hanxinzhi", feishu_notice.UserId13)
		//feishu_notice.CandyNotice("145hanxinzhi", feishu_notice.UserId13)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chenkaige", feishu_notice.UserId14)
		//feishu_notice.CandyNotice("chenkaige", feishu_notice.UserId14)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("xiaxueying", feishu_notice.UserId16)
		//feishu_notice.CandyNotice("xiaxueying", feishu_notice.UserId16)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145xiaoxiaohao", feishu_notice.UserId15)
		//feishu_notice.CandyNotice("145xiaoxiaohao", feishu_notice.UserId15)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("fadianxiaohao", feishu_notice.UserId18)
		//feishu_notice.CandyNotice("fadianxiaohao", feishu_notice.UserId18)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("jerry", feishu_notice.UserId19)
		//feishu_notice.CandyNotice("jerry", feishu_notice.UserId19)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("bingqiling", feishu_notice.UserId20)
		//feishu_notice.CandyNotice("bingqiling", feishu_notice.UserId20)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("xiyixiaohao", feishu_notice.UserId21)
		//feishu_notice.CandyNotice("xiyixiaohao", feishu_notice.UserId21)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("tiedan", feishu_notice.UserId22)
		//feishu_notice.CandyNotice("tiedan", feishu_notice.UserId22)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("mentongxiaohao", feishu_notice.UserId23)
		//feishu_notice.CandyNotice("mentongxiaohao", feishu_notice.UserId23)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("myrxh", feishu_notice.UserId24)
		//feishu_notice.CandyNotice("myrxh", feishu_notice.UserId24)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("wxf", feishu_notice.UserId25)
		//feishu_notice.CandyNotice("wxf", feishu_notice.UserId25)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("peizhixiaohao", feishu_notice.UserId27)
		//feishu_notice.CandyNotice("peizhixiaohao", feishu_notice.UserId27)
	}); err != nil {
		panic(err)
	}

	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("test", feishu_notice.UserId26)
	//}); err != nil {
	//	panic(err)
	//}

	c.Start()
}
