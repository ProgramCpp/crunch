package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/buger/jsonparser"
)

type counterHandler struct {
	counter *counter
}

func NewCounterHandler(c *counter) counterHandler {
	return counterHandler{
		counter: c,
	}
}

func (c counterHandler)Handle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	count, err := jsonparser.GetInt(data, "count")
	if err != nil {
		log.Printf("error reading count: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	count = atomic.AddInt64(&c.counter.value, count)

	w.Write([]byte( fmt.Sprintf(
	`
	{
		"count": %d
	}
	`, count)))
}
