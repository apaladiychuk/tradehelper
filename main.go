package main

import "fmt"

var Config AppConfig

func main() {
	if err := LoadConfig(); err != nil {
		fmt.Errorf(" load config %v ", err)
		return
	}

	apis := InitApi()

	for _, api := range apis {
		go api.Schedule()
	}
	m := make(chan int)
	ScheduleRequest(&m, 0)

	res := <-m
	fmt.Printf(" chann %d  \n", res)
}
