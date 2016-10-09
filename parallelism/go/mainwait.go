package main

import (
	"flag"
	"fmt"
	"sync"
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

func main() {
	taskCount := flag.Int("tasks", 1, "count of task to parallelize")
	flag.Parse()

	if *taskCount == 0 {
		flag.Usage()
		return
	}

	fmt.Printf("Task to execute: %d", *taskCount)
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(*taskCount)

	start := time.Now()
	result := make([]uint32, *taskCount)

	for i := 0; i < *taskCount; i++ {
		go func(i int, r *uint32) {
			s := fmt.Sprintf("Task %d done!", i)
			time.Sleep(100 * time.Microsecond)
			*r = hash(s)
			wg.Done()
		}(i, &result[i])
	}

	wg.Wait()

	h := uint32(0)
	for i := 0; i < *taskCount; i++ {
		h ^= result[i]
	}
	elapsed := time.Since(start)
	fmt.Printf("%d in %s, hash = 0x%x", *taskCount, elapsed, h)
	fmt.Println()
}
