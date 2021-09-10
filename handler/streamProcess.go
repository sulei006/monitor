package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"monitor/dal"
	"monitor/model"
	"monitor/utils"
	"net/http"
	"strconv"
)

//这里后面要处理一下，感觉不应该使用fatal
func GetStreamData(context *gin.Context) {
	startHour, err := strconv.Atoi(context.PostForm("startHour"))
	if err != nil {
		log.Printf("illegal param 'startHour: %s'", context.PostForm("startHour"))
	}
	startMin, err := strconv.Atoi(context.PostForm("startMin"))
	if err != nil {
		log.Printf("illegal param 'startMin: %s'", context.PostForm("startMin"))
	}
	endHour, err := strconv.Atoi(context.PostForm("endHour"))
	if err != nil {
		log.Printf("illegal param 'endHour: %s'", context.PostForm("endHour"))
	}
	endMin, err := strconv.Atoi(context.PostForm("endMin"))
	if err != nil {
		log.Printf("illegal param 'endMin: %s'", context.PostForm("endMin"))
	}
	lower := startHour*100 + startMin
	upper := endHour*100 + endMin
	streamData, err := dal.GetSteamDataFromDb(lower, upper)
	if err != nil {
		log.Printf("get stream data from db error: %s", err)
	}
	streamDataVo := []model.StremDataVo{}

	for _, sdata := range streamData {
		sdataVo := model.StremDataVo{}
		sdataVo.Input = utils.GetStreamData(sdata.Input)
		sdataVo.Output = utils.GetStreamData(sdata.Output)
		sdataVo.Id = sdata.Id
		streamDataVo = append(streamDataVo, sdataVo)
	}
	context.JSON(http.StatusOK, gin.H{
		"stream-data": streamDataVo,
	})
}
