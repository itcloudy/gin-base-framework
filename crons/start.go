package crons

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/robfig/cron"
)

//const companyUpdateTimmer = "@monthly"
//const updateScoreTimmer = "@daily"
//
////公司修改次数，每月更新
//func companyUpdatetimer() {
//	var modelArray []models.Company
//	var CompanyEditNums []models.CompanyEditNums
//	common.DB.Find(&modelArray)
//	common.DB.Find(&CompanyEditNums)
//	for _, v := range modelArray {
//		for _, v1 := range CompanyEditNums {
//
//		}
//	}
//}
//初始化当天统计数据
func initTodayDataNum() {
	client := common.RedisClient
	client.Del(common.TODAY_COMPANY_APPLY_TOTAL)
	client.Del(common.TODAY_IP_ARRAY)
	client.Del(common.TODAY_VIEWS)
}


func Run() {
	c := cron.New()


	c.Start()
	select {}
}
