package main

import "fmt"

var Config AppConfig

func main() {
	if err := LoadConfig(); err != nil {
		fmt.Errorf(" load config %v ", err)
		return
	}

	apis := InitApi()
	m := make(chan int)
	for _, api := range apis {
		go api.Schedule(&m)
	}
	res := <-m
	fmt.Printf(" chann %d  \n", res)
}
