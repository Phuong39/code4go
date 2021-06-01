package crontab

import "github.com/robfig/cron"

var crontab = cron.New()

// RegisterCrontab 定时任务注册
func RegisterCrontab(spec string, f func()) error {
	return crontab.AddFunc(spec, f)
}

// CrontabRun 定时任务启动
func CrontabRun() {
	crontab.Run()
}
