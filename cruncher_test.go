package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/buger/jsonparser"

	"github.com/stretchr/testify/assert"
)

func TestCounterHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/",
		strings.NewReader(
		`{
			"count": 10 
		}`),
	)
	assert.NoError(t, err)

	res := httptest.NewRecorder()
	counterHandler(res, req)

	assert.Equal(t, http.StatusOK,res.Code);

	data, err := io.ReadAll(res.Result().Body)
	assert.NoError(t, err)

	count, err :=jsonparser.GetInt(data, "count")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count);
}


