package main

import (
	"salias/service/api"
	"log"
	"net/http"
)

func main() {
	port := "9020"
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":" + port, router))
}