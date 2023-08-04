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
		feishu_notice.MotorNotice(feishu_notice.Token1, "大佬")
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token2, "富豪")
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token3, "145")
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token4, "谜语人")
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token5, "柴郡猫")
	}); err != nil {
		panic(err)
	}

	if err := c.AddFunc("*/6 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token6, "配置")
	}); err != nil {
		panic(err)
	}

	//if err := c.AddFunc("*/6 * * * * *", func() {
	//	feishu_notice.MotorNotice(feishu_notice.Token7, "陆逊")
	//}); err != nil {
	//	panic(err)
	//}

	//if err := c.AddFunc("*/5 * * * * *", func() {
	//	feishu_notice.MotorNotice(feishu_notice.Token8, "145胡莹")
	//}); err != nil {
	//	panic(err)
	//}

	if err := c.AddFunc("*/5 * * * * *", func() {
		feishu_notice.MotorNotice(feishu_notice.Token9, "门童")
	}); err != nil {
		panic(err)
	}
	c.Start()
}
