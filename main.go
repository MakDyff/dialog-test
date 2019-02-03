package main

import (
	"Makdyff/dialog-test/services"
	"encoding/json"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	pingS := services.NewPingService()
	fileS := services.NewFileService("conf/sites.txt")
	statS := services.NewStatisticService()

	pingS.AddHandler(statS.Ping)
	fileS.AddHandler(pingS.UrlString)

	pingS.Run()
	fileS.Run()
	statS.Run()

	logs.Info("Run test...")

	go tmp1(statS)
	go tmp2(statS)
	go tmp3(statS)

	var gracefulStop = make(chan os.Signal)
	<-gracefulStop

	logs.Warn("Exit")
}

func tmp1(stat *services.StatisticService) {
	for {
		time.Sleep(5 * time.Second)
		t := stat.GetInfo("xvideos.com")
		tt, _ := json.Marshal(t)
		logs.Warn(string(tt))
	}
}

func tmp2(stat *services.StatisticService) {
	for {
		time.Sleep(20 * time.Second)
		t := stat.GetMaxMinPing()

		tt, _ := json.Marshal(t)
		logs.Warn(string(tt))
	}
}

func tmp3(stat *services.StatisticService) {
	for {
		time.Sleep(20 * time.Second)
		t := stat.RequestStat()

		tt, _ := json.Marshal(t)
		logs.Warn(string(tt))
	}
}
