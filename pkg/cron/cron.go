package cron

import (
	"log"

	"github.com/laenix/axhub/pkg/axhub"
	"github.com/laenix/axhub/pkg/lark"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func EveryMinute() {
	c := cron.New()
	// do query every 30 seconds
	c.AddFunc("*/30 * * * *", func() {
		url := viper.GetString("app.url.axhub")
		target, time := axhub.GetAxhubInfo(url)
		apiurl := viper.GetString("app.url.lark")
		if viper.GetString("app.timestamp") != time {
			log.Println("[+]", target, time, url, apiurl)
			// Write a new timestamp into config file
			viper.GetViper().Set("app.timestamp", time)
			viper.WriteConfig()
			lark.SendMessage(apiurl, url, target, time)
		}
	})
	c.Start()
	select {}
}
