package main

import (
	"log"
	"net/http"
)

func main(){
	c := NewCounter()
	c.Run()

	http.HandleFunc("/", NewCounterHandler(&c).Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}