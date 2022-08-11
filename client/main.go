package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*10)
				if err != nil {
					return nil, err
				}
				//c.SetDeadline(time.Now().Add(time.Second * 5))
				return c, nil
				//return c, nil
			},
			MaxIdleConnsPerHost: 500,
			DisableKeepAlives:   false,
			//not set response time out
			//ResponseHeaderTimeout: time.Second * 10,
		},
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:9999/hello", nil)
	if err != nil {
		return
	}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	if rsp.StatusCode != 200 {
		fmt.Println("status err", rsp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("read body err")
		return
	}
	rsp.Body.Close()
	fmt.Println("body ", string(body))
}
