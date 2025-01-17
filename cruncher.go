package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/buger/jsonparser"
)

// assuming the max counter value is in range.
// todo: add wraparound for larger values
var counter int64

func counterHandler(w http.ResponseWriter, r *http.Request) {
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

	count = atomic.AddInt64(&counter, count)

	w.Write([]byte( fmt.Sprintf(
	`
	{
		"count": %d
	}
	`, count)))
}
