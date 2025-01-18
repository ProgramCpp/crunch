package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	c.Run()

	c.Add(10)
	c.Add(10)
	c.Add(10)

	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, int64(30), c.Value())
}