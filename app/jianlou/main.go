package main

import (
	"notice/app/jianlou/api"
	"notice/internal/cli"
)

func init() {
	cli.InitRedisClient()
}
func main() {
	api.Begin()
}
