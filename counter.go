package main

type counter struct {
	// assuming the max counter value is in range.
	// todo: add wraparound for larger values
	value int64
}
