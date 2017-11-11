package main

import (
	"net/http"
	"log"
	"github.com/ThoreauZZ/go-restful-api/routers"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        routers.UserResource(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
