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
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145", feishu_notice.UserId3)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("miyuren", feishu_notice.UserId4)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chaijunmao", feishu_notice.UserId5)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("peizhi", feishu_notice.UserId6)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145胡莹", feishu_notice.UserId7)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("门童", feishu_notice.UserId8)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("富豪小号", feishu_notice.UserId9)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("xiyijueyuanti", feishu_notice.UserId10)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("fadian", feishu_notice.UserId11)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145hanxinzhi", feishu_notice.UserId13)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chenkaige", feishu_notice.UserId14)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("xiaxueying", feishu_notice.UserId16)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145xiaoxiaohao", feishu_notice.UserId15)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("fadianxiaohao", feishu_notice.UserId18)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("jerry", feishu_notice.UserId19)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("bingqiling", feishu_notice.UserId20)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("xiyixiaohao", feishu_notice.UserId21)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("tiedan", feishu_notice.UserId22)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("mentongxiaohao", feishu_notice.UserId23)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("myrxh", feishu_notice.UserId24)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("wxf", feishu_notice.UserId25)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("peizhixiaohao", feishu_notice.UserId27)
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
