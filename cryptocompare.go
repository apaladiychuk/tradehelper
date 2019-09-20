package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CryptoCompare struct{
	Config Apiconfig
}

func( c *CryptoCompare) Schedule(m  *chan int){
	fmt.Println("start timer ")
	timer := time.NewTimer(SCHEDULER_TIME)
	if ii == 10 {
		*m  <- ii
		return
	}
	ii ++
	go func(timer *time.Timer) {
		fmt.Println("wait  timer ")
		tmr := <-timer.C
		fmt.Println("run  timer ", tmr)
		c.requestCrypto()
		c.Schedule(m, )
	}(timer)

}

func( c *CryptoCompare)  requestCrypto() {
	client := http.Client{}
	url := c.Config.Url + c.Config.
	request, err := http.NewRequest("GET","https://min-api.cryptocompare.com/data/v2/histoday?fsym=BTC&tsym=USD&limit=10",bytes.NewBufferString(""))
	if err != nil {
		fmt.Errorf("err>> %v",err )
		return
	}
	request.Header.Add("authorization", )
	if resp , err := client.Do(request) ;err != nil{
		fmt.Errorf("Responce Err >> %v ", err)
	} else {
		defer resp.Body.Close()
		if body, err  := ioutil.ReadAll(resp.Body) ; err != nil{
			fmt.Errorf("parse body %v ", err)
		} else {
			fmt.Printf(" [OK]  %v \n",body )
			fmt.Println(" [OK]  --- ",string(body ))
		}
	}
}