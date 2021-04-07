package main

import (
	"fmt"
	"time"

	"github.com/lack-io/gscheduler"
)

func main() {
	s := gscheduler.Default()
	s.Start()

	a := 1
	j, _ := s.AddIntervalJob("job1", true, time.Second*1, func() {
		fmt.Printf("[%s] a = %d\n", time.Now(), a)
		a++
	})

	s.AddIntervalJob("job2", true, time.Second*2, func() {
		println("job2")
	})

	time.Sleep(time.Second * 3)
	_ = j
	s.StopGraceful()

	time.Sleep(time.Second * 4)
}