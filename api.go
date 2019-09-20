package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const SCHEDULER_TIME = 5 * time.Second

func ScheduleRequest(m *chan int, ii int) {

	fmt.Println("start timer ")
	timer := time.NewTimer(SCHEDULER_TIME)
	if ii == 10 {
		*m <- ii
		return
	}
	ii++
	go func(timer *time.Timer) {
		fmt.Println("wait  timer ")
		tmr := <-timer.C
		fmt.Println("run  timer ", tmr)
		requestCrypto()
		ScheduleRequest(m, ii)
	}(timer)
}
