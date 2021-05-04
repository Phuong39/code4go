package times

import "github.com/robfig/cron"

func NewTimingJob(spec string, cmd func()) error {
	c := cron.New()
	err := c.AddFunc(spec, func() {
		cmd()
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
