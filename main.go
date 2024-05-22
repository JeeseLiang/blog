package main

import (
	"blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   10 * time.Second,
	}
	s.ListenAndServe()
}
