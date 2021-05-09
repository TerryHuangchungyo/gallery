package main

import (
	"gallery-frontend/config"
	"gallery-frontend/controller"
	. "gallery-frontend/utils"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:           config.App.Port,
		Handler:        controller.Mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		ErrorLog:       ErrorLog,
		MaxHeaderBytes: 1 << 20,
	}

	InfoLog.Println("Server Start......")
	if err := server.ListenAndServe(); err != nil {
		ErrorLog.Printf("Http server error: %v", err)
		panic(err)
	}
}
