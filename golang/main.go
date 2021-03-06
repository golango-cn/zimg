package main

import (
	"fmt"
	"net/http"
	"time"
	"zimg/core/common"
	routers "zimg/router"
)

func main()  {

	common.Init()

	//运行模式
	routersInit := routers.InitRouter()

	port := fmt.Sprintf(":%d", common.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           port,
		Handler:        routersInit,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()

}
