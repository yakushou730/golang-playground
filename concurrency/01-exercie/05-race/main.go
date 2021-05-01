package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	var t *time.Timer
	ch := make(chan bool)
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true

	})
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}

// returns random duration between 0-1 seconds
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
