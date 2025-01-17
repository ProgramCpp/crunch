package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", counterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}