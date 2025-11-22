package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var db = []string{"A", "B", "C", "D"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(db); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v\n", results)
}

func dbCall(i int) {
	// simulate db call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", db[i])
	m.Lock()
	results = append(results, db[i])
	m.Unlock()
	wg.Done()
}
