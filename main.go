package main

import (
	"fmt"
	"net/http"
	"z-img/config"
	_ "z-img/database"
	"z-img/router"
)

func main() {
	r := router.InitRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HttpPort),
		Handler:      r,
		ReadTimeout:  config.ReadTimeOut,
		WriteTimeout: config.WriteTimeOut,
	}

	_ = server.ListenAndServe()
}
