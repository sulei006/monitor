package controller

import (
	"github.com/gin-gonic/gin"
	"monitor/handler"
)

func StreamDataQueryEntrance() {
	router := gin.Default()
	router.POST("/stream",handler.GetStreamData)
	router.Run()
}
