package main

import (
	"net/http"
	"time"

	"github.com/th997/goproxy-tls/goproxy"
)

func main() {
	proxy := goproxy.New()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      proxy,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
