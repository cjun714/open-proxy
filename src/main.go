package main

import (
	"net/http"
	"util/log"

	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/open/token", token)
	router.GET("/poc/hello", hello)
	router.NotFound = http.FileServer(http.Dir("page"))

	log.H("Server is started")

	log.E(http.ListenAndServe(":"+os.Args[1], router))
}
