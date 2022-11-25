package main

import (
	"log"

	"github.com/laenix/axhub/config"
	"github.com/laenix/axhub/pkg/cron"
)

func main() {
	log.Println("[+] start axhub notice to lark")
	config.InitConfig()
	cron.EveryMinute()
}
