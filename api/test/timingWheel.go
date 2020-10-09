package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/collection"
	"time"
)

func main() {
	timingWheel, err := collection.NewTimingWheel(time.Second, 300, func(k, v interface{}) {
		fmt.Println(k, v)
	})
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		timingWheel.SetTimer(i, i, time.Second*time.Duration(i))
	}
	for i := 1; i <= 10; i++ {
		if i >= 5 {
			timingWheel.MoveTimer(i, time.Duration(i*2)*time.Second)
		}
	}
	time.Sleep(time.Minute)
}
