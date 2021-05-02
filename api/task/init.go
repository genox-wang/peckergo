package task

import (
	"github.com/robfig/cron"
)

const (
	// EveryHour 每小时
	EveryHour = "0 0 * * * *"
)

var (
	c *cron.Cron
)

func init() {
	c = cron.New()
	// RegisterTask(collectPVUV, EveryHour)
	// RegisterTask(collectReqCost, EveryHour)
	c.Start()
}

// RegisterTask Register task
func RegisterTask(cb func(), duration string) {
	c.AddFunc(duration, cb)
}
