package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func healthCheckAllServers() {
	cron := gocron.NewScheduler(time.Local)
	cron.Every(2).Seconds().Do(func() {
		for i := range servers {
			servers[i].CheckHealth()
		}
	})
	<-cron.StartAsync()
}
