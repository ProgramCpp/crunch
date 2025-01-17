package main

import (
	"log"
	"net/http"
)

func main(){
	var c counter

	http.HandleFunc("/", NewCounterHandler(&c).Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}