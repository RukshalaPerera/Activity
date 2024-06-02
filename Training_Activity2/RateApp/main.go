package main

import (
	"fmt"
	"time"

	"RateApp/app/handler"
	"RateApp/app/model"
)

func main() {
	tb := model.NewTokenBucket(10, 1)
	for i := 0; i < 20; i++ {
		fmt.Printf("Request %d: %v\n", i+1, handler.Request(tb, 1))
		time.Sleep(500 * time.Millisecond)
	}
}
