package main

import (
	"github.com/serkas/salias/service/api"
	"log"
	"net/http"
)

func main() {
	port := "9020"
	router := api.NewRouter()
	log.Printf("Now start server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
