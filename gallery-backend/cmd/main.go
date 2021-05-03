package main

import (
	"gallery-backend/controller"
	. "gallery-backend/utils"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:           ":8080",
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
