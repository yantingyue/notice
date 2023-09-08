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

	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("富豪", feishu_notice.UserId2)
	//}); err != nil {
	//	panic(err)
	//}

	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("145", feishu_notice.UserId3)
	//}); err != nil {
	//	panic(err)
	//}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("谜语人", feishu_notice.UserId4)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("柴郡猫", feishu_notice.UserId5)
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice("配置", feishu_notice.UserId6)
	}); err != nil {
		panic(err)
	}

	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("145胡莹", feishu_notice.UserId7)
	//}); err != nil {
	//	panic(err)
	//}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("门童", feishu_notice.UserId8)
	}); err != nil {
		panic(err)
	}
	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("富豪小号", feishu_notice.UserId9)
	//}); err != nil {
	//	panic(err)
	//}

	//if err := c.AddFunc("*/7 * * * * *", func() {
	//	feishu_notice.MotorNotice("蜥蜴绝缘体", feishu_notice.UserId10)
	//}); err != nil {
	//	panic(err)
	//}
	//if err := c.AddFunc("*/7 * * * * *", func() {
	//	feishu_notice.MotorNotice("法典人", feishu_notice.UserId11)
	//}); err != nil {
	//	panic(err)
	//}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145韩新枝", feishu_notice.UserId13)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("陈凯歌", feishu_notice.UserId14)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("145小号", feishu_notice.UserId15)
	}); err != nil {
		panic(err)
	}
	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice("夏雪莹", feishu_notice.UserId16)
	}); err != nil {
		panic(err)
	}
	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice("水镜先生", feishu_notice.UserId17)
	//}); err != nil {
	//	panic(err)
	//}
	//if err := c.AddFunc("*/7 * * * * *", func() {
	//	feishu_notice.MotorNotice("test", feishu_notice.UserId12)
	//}); err != nil {
	//	panic(err)
	//}
	c.Start()
}
