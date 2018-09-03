package main

import (
	go_log "log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"be/common/log"
	"be/handle"
	"be/mysql"
	"be/options"
	"be/server"
)

func doServe() {
	defer func() {
		if err := recover(); err != nil {
			doServe()
		}
	}()

	// 初始化DB
	mysql.DB.InitConn()
	// 初始化服务,并启动服务
	mux := server.New()
	// URL映射
	handle.InitHandle(mux)
	srv := &http.Server{
		Handler:      mux.GetRouter(),
		Addr:         options.Options.HTTPAddress + ":" + strconv.FormatUint(options.Options.HTTPPort, 10),
		WriteTimeout: 15 * time.Hour,
		ReadTimeout:  15 * time.Hour,
		ErrorLog:     go_log.New(log.StandardLogger().Writer(), "", 0),
	}

	// 启动主服务
	log.Fatal(srv.ListenAndServe())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	// 从命令行、配置文件初始化配置
	options.Options.InitOptions()

	// 初始化Log
	log.InitLog()
	// 可以使用log了
	log.Infoln("日志文件初始化成功")
	// 启动服务
	doServe()
}
