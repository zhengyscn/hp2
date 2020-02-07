package main

import (
	"fmt"
	"hp2/config"
	"hp2/dao"
	"hp2/routers"
	"net/http"
	"os"
)

func init() {
	var err error
	// Set up database connection amd sync table
	dao.Setup()
	dao.AutoMigrate()

	// Set rest auth
	err = dao.NewCasbin()
	if err != nil {
		fmt.Printf("NewCasbin err: %v\n", err)
		os.Exit(-1)
	}
}

func main() {
	router := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", config.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Printf("[info] start http server listening %s\n", endPoint)
	err := server.ListenAndServe()
	fmt.Println(err)
}
