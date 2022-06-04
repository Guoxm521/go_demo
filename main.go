package main

import (
	"example.com/m/v2/service/spider_service"
)

func main() {
	//router := routers.InitRouter()
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	//	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}

func init() {
	//setting.Setup()
	//models.Setup()

	spider_service.Baidu()
}
