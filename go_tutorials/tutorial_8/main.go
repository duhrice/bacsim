package main

import (
	"fmt"
	"time"
	"sync"
)

// Create a mutex to prevent tasks from wusing the same important resource at once
var m = sync.RWMutex {}
// Create a wait group to wait for tasks that are running concurrently
var wg = sync.WaitGroup {}
var dbData = []string {"id1", "id2", "id3", "id4", "id5"}
var results =  []string {}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	// Notify wait group that task has finished running, decrementing wait group counter
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// Wait group acts essentially like a counter for the number of task running concurrently
		wg.Add(1)
		// Enable concurrency thorugh parallel execution by using "go" keyword (go rountines)
		go dbCall(i)
	}
	// Wait for all tasks to finish in wait group
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v\n", results)
	// go run .\go_tutorials\tutorial_8\main.go
}