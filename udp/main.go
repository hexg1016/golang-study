package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Action(tmpI int) {

	addr, err := net.ResolveIPAddr("ip", "127.0.0.1")
	if err != nil {
		fmt.Println("Resolvtion error", err.Error())

	}

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   addr.IP,
		Port: 30011,
	})
	if err != nil {
		log.Println("Connect to udp server failed,err:", err)
		return
	}

	for i := 0; i < 10000; i++ {
		//startTm := time.Now().UnixNano()
		_, err := conn.Write([]byte(fmt.Sprintf("%v", i)))
		if err != nil {
			fmt.Printf("Send data failed,err:", err)
			return
		}

		//接收数据
		result := make([]byte, 64)
		//middleTm := time.Now().UnixNano()
		_, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("Read from udp server failed ,err:", err)
			return
		}
		//end := time.Now().UnixNano()
		fmt.Printf("Recived msg from %s,  %d  \n", remoteAddr, tmpI)
		//time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	startTm := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmpI := i
		//time.Sleep(1 * time.Second)
		go Action(tmpI)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("end %d \n", end-startTm)

}
