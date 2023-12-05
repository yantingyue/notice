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
		feishu_notice.CandyNotice("fuhaotangguo", feishu_notice.UserId2)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145", feishu_notice.UserId3)
		feishu_notice.CandyNotice("145tangguo", feishu_notice.UserId3)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("miyuren", feishu_notice.UserId4)
		feishu_notice.CandyNotice("miyurentangguo", feishu_notice.UserId4)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chaijunmao", feishu_notice.UserId5)
		feishu_notice.CandyNotice("chaijunmaotangguo", feishu_notice.UserId5)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("peizhi", feishu_notice.UserId6)
		feishu_notice.CandyNotice("peizhitangguo", feishu_notice.UserId6)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145胡莹", feishu_notice.UserId7)
		feishu_notice.CandyNotice("145胡莹糖果", feishu_notice.UserId7)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("门童", feishu_notice.UserId8)
		feishu_notice.CandyNotice("门童糖果", feishu_notice.UserId8)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("富豪小号", feishu_notice.UserId9)
		feishu_notice.CandyNotice("富豪小号糖果", feishu_notice.UserId9)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("xiyijueyuanti", feishu_notice.UserId10)
		feishu_notice.CandyNotice("xiyijueyuantitangguo", feishu_notice.UserId10)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/7 * * * * *", func() {
		feishu_notice.MotorNotice("fadian", feishu_notice.UserId11)
		feishu_notice.CandyNotice("fadiantangguo", feishu_notice.UserId11)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145hanxinzhi", feishu_notice.UserId13)
		feishu_notice.CandyNotice("145hanxinzhitangguo", feishu_notice.UserId13)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("chenkaige", feishu_notice.UserId14)
		feishu_notice.CandyNotice("chenkaigetangguo", feishu_notice.UserId14)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("xiaxueying", feishu_notice.UserId16)
		feishu_notice.CandyNotice("xiaxueyingtangguo", feishu_notice.UserId16)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145xiaoxiaohao", feishu_notice.UserId15)
		feishu_notice.CandyNotice("145xiaoxiaohaotangguo", feishu_notice.UserId15)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("fadianxiaohao", feishu_notice.UserId18)
		feishu_notice.CandyNotice("fadianxiaohaotangguo", feishu_notice.UserId18)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("jerry", feishu_notice.UserId19)
		feishu_notice.CandyNotice("jerrytangguo", feishu_notice.UserId19)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("bingqiling", feishu_notice.UserId20)
		feishu_notice.CandyNotice("bingqilingtangguo", feishu_notice.UserId20)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("xiyixiaohao", feishu_notice.UserId21)
		feishu_notice.CandyNotice("xiyixiaohaotangguo", feishu_notice.UserId21)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("tiedan", feishu_notice.UserId22)
		feishu_notice.CandyNotice("tiedantangguo", feishu_notice.UserId22)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("mentongxiaohao", feishu_notice.UserId23)
		feishu_notice.CandyNotice("mentongxiaohaotangguo", feishu_notice.UserId23)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("myrxh", feishu_notice.UserId24)
		feishu_notice.CandyNotice("myrxhtangguo", feishu_notice.UserId24)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("wxf", feishu_notice.UserId25)
		feishu_notice.CandyNotice("wxftangguo", feishu_notice.UserId25)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("peizhixiaohao", feishu_notice.UserId27)
		feishu_notice.CandyNotice("peizhixiaohaotangguo", feishu_notice.UserId27)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("xlz", feishu_notice.UserId28)
		feishu_notice.CandyNotice("xlztangguo", feishu_notice.UserId28)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("shj", feishu_notice.UserId29)
		feishu_notice.CandyNotice("shjtangguo", feishu_notice.UserId29)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("test", feishu_notice.UserId26)
	}); err != nil {
		panic(err)
	}

	//if err := c.AddFunc("*/15 * * * * *", func() {
	//	feishu_notice.NiceNotice("wEaOgGs2ulepxrsMlvimPoQSMxE3r3HO")
	//}); err != nil {
	//	panic(err)
	//}
	//if err := c.AddFunc("*/20 * * * * *", func() {
	//	feishu_notice.NiceNotice("MwGcQTsTGXTmPwDmPvTnPXiNQkwc-Ar6")
	//}); err != nil {
	//	panic(err)
	//}
	//if err := c.AddFunc("*/10 * * * * *", func() {
	//	feishu_notice.DTList("MwGcQTsTGXTmPwDmPvTnPXiNQkwc-Ar6")
	//}); err != nil {
	//	panic(err)
	//}
	c.Start()
}
