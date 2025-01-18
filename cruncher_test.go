package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCounterHandler(t *testing.T) {
	c := NewCounter()
	c.Run()
	handler := NewCounterHandler(&c)

	req, err := http.NewRequest(http.MethodGet, "/",
		strings.NewReader(`{
			"count": 10 
		}`),
	)
	assert.NoError(t, err)

	res := httptest.NewRecorder()

	handler.Handle(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, int64(10), c.Value())
}

func TestCounterHandlerMultipleRequests(t *testing.T) {
	c := NewCounter()
	c.Run()
	handler := NewCounterHandler(&c)

	for i := 1; i <= 10; i++ {
		req, err := http.NewRequest(http.MethodGet, "/",
			strings.NewReader(`{
				"count": 10 
			}`),
		)
		assert.NoError(t, err)

		res := httptest.NewRecorder()

		handler.Handle(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	}

	assert.Equal(t, int64(100), c.Value()) // counter updated after all the api calls
}

func TestCounterHandlerConcurrentRequests(t *testing.T) {
	c := NewCounter()
	c.Run()
	handler := NewCounterHandler(&c)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, "/",
				strings.NewReader(`{
				"count": 10 
			}`),
			)
			assert.NoError(t, err)

			res := httptest.NewRecorder()

			handler.Handle(res, req)
			assert.Equal(t, http.StatusOK, res.Code)
		}()
	}

	wg.Wait()

	time.Sleep(1 * time.Second)

	assert.Equal(t, int64(10000), c.Value())
}
