package main

import "sync/atomic"

type counter struct {
	// assuming the max counter value is in range.
	// todo: add wraparound for larger values
	value int64

	stream chan int64
}

func NewCounter() counter {
	return counter{
		stream: make(chan int64, 10000), // max delay in value updates.
	}
}

func (c counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func (c *counter) Add(x int64) {
	c.stream <- x
}

// todo: error handling
func (c *counter) Run() {
	go func() {
		for v := range c.stream {
			atomic.AddInt64(&c.value, v)
		}
	}()
}
