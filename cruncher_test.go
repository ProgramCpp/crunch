package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/buger/jsonparser"

	"github.com/stretchr/testify/assert"
)

func TestCounterHandler(t *testing.T) {
	var c counter
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

	data, err := io.ReadAll(res.Result().Body)
	assert.NoError(t, err)

	count, err := jsonparser.GetInt(data, "count")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)
}

func TestCounterHandlerMultipleRequests(t *testing.T) {
	var c counter
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

		data, err := io.ReadAll(res.Result().Body)
		assert.NoError(t, err)

		count, err := jsonparser.GetInt(data, "count")
		assert.NoError(t, err)
		assert.Equal(t, int64(i * 10), count) // counter updated after each api call
	}

	assert.Equal(t, int64(100), c.value) // counter updated after all the api calls
}

func TestCounterHandlerConcurrentRequests(t *testing.T) {
	var c counter
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

	assert.Equal(t, int64(10000), c.value)
}
