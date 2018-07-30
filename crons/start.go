package crons

import (
	"github.com/robfig/cron"
)


func updateScore() {

}



func Run() {
	c := cron.New()
	//每天1点更新
	c.AddFunc("0 0 1 * * ?", updateScore)

	c.Start()
	select {}
}
