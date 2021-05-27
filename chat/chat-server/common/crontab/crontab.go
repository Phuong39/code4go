package crontab

import "github.com/robfig/cron"

var crontab = cron.New()

func RegisterCrontab(spec string, f func()) error {
	return crontab.AddFunc(spec, f)
}

func CrontabRun() {
	crontab.Run()
}
