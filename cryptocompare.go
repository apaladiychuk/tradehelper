package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const SCHEDULER_TIME = 5 * time.Second

type CryptoCompare struct {
	Config Apiconfig
}

func (c *CryptoCompare) Schedule(m *chan int) {
	fmt.Println("start timer ")
	timer := time.NewTimer(SCHEDULER_TIME)
	go func(timer *time.Timer) {
		fmt.Println("wait  timer ")
		tmr := <-timer.C
		fmt.Println("run  timer ", tmr)
		c.requestCrypto()
		c.Schedule(m)
	}(timer)

}

func (c *CryptoCompare) requestCrypto() {
	client := http.Client{}
	r := c.Config.Param["range"]
	fsym := c.Config.Param["fsym"]
	tsym := c.Config.Param["tsym"]
	limit := c.Config.Param["limit"]
	url := fmt.Sprintf(c.Config.Url+"%s?fsym=%s&tsym=%s&limit=%s", r, fsym, tsym, limit)
	request, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
	if err != nil {
		fmt.Errorf("err>> %v", err)
		return
	}
	request.Header.Add("authorization", c.Config.ApiKey)
	if resp, err := client.Do(request); err != nil {
		fmt.Errorf("Responce Err >> %v ", err)
	} else {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Errorf("parse body %v ", err)
		} else {
			var data CryptoCompareResponse
			if err := json.Unmarshal(body, &data); err != nil {
				fmt.Println(" Json unmarshal error ", err.Error())
			} else {
				fmt.Printf(" Json resp %v \n ", data)
			}
			//fmt.Println(" [OK]  --- ",string(body ))
		}
	}
}
