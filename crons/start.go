package crons

import (
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/robfig/cron"
)

func initTodayDataNum() {

}

func Run() {
	c := cron.New()
	// two method
	c.AddFunc(common.Quartz.UpdateScoreTime, initTodayDataNum)
	c.AddFunc("0 0 1 * * ?", initTodayDataNum)
	c.Start()
	select {}
}
