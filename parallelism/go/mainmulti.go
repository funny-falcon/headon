package main

import (
	"flag"
	"fmt"
	"time"
)

func hash(s string) uint32 {
	h := uint32(0xcafe)
	l := len(s)
	for i := 0; i < l; i++ {
		h ^= uint32(s[i])
		h *= 0x11f
	}
	return h
}

func work(i int, ch chan uint32) {
	s := fmt.Sprintf("Task %d done!", i)
	time.Sleep(100 * time.Microsecond)
	ch <- hash(s)
}

func main() {
	taskCount := flag.Int("tasks", 1, "count of task to parallelize")
	flag.Parse()

	if *taskCount == 0 {
		flag.Usage()
		return
	}

	fmt.Printf("Task to execute: %d", *taskCount)
	fmt.Println()

	start := time.Now()
	results := make([]chan uint32, *taskCount)

	for i := 0; i < *taskCount; i++ {
		ch := make(chan uint32, 1)
		go work(i, ch)
		results[i] = ch
	}

	h := uint32(0)
	for i := 0; i < *taskCount; i++ {
		h ^= <-results[i]
	}
	elapsed := time.Since(start)
	fmt.Printf("%d in %s, hash = 0x%x", *taskCount, elapsed, h)
	fmt.Println()
}
