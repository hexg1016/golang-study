package main

import (
	"fmt"
	"time"
)

var SecondCheckTimer1 *time.Timer
var SecondCheckTimer2 *time.Timer
var SecondCheckTimer3 *time.Timer
var cnt1, cnt2, cnt3 int

func main() {
	SecondCheckTimer1 = time.NewTimer(1 * time.Second)
	SecondCheckTimer2 = time.NewTimer(1 * time.Second)
	SecondCheckTimer3 = time.NewTimer(1 * time.Second)
	timeBegin := int32(time.Now().Unix())
	for {
		select {
		case <-SecondCheckTimer1.C:
			cnt1++
			fmt.Println("cnt1 : ", cnt1, "time: ", int32(time.Now().Unix())-timeBegin)
			time.Sleep(1 * time.Second)
			SecondCheckTimer1.Reset(1 * time.Second)
		case <-SecondCheckTimer2.C:
			cnt2++
			fmt.Println("cnt2 : ", cnt2, "time: ", int32(time.Now().Unix())-timeBegin)
			time.Sleep(1 * time.Second)
			SecondCheckTimer2.Reset(1 * time.Second)
		case <-SecondCheckTimer3.C:
			cnt3++
			fmt.Println("cnt3 : ", cnt3, "time: ", int32(time.Now().Unix())-timeBegin)
			time.Sleep(1 * time.Second)
			SecondCheckTimer3.Reset(1 * time.Second)
		}
	}
}
