package main

import (
	"monitor/controller"
	"monitor/dal"
	"monitor/utils"
)

func main(){
	c := utils.GetConf("./dev.yaml")
	dal.DatabaseConnect(c.DbHost,c.DbPort,c.DbUser,c.DbPwd,c.DbDataBase)
	controller.StreamDataQueryEntrance()


}


