package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:10080")
	if err != nil {
		fmt.Printf("udp resolve fail: err %v", err)
	}
	raddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:10086")
	if err != nil {
		fmt.Printf("udp resolve fail: err %v", err)
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Printf("udp listen fail: err %v", err)
	}

	go Op(conn)

	for {
		time.Sleep(time.Second)
	}
	
}

func Op(conn *net.UDPConn) {
	for {
	   msg := "x"
	   buf := []byte(msg)
	   _, err := conn.Write(buf)
	   if err != nil {
         fmt.Printf("write fail: err %v", err)
	   }
	   fmt.Printf("[write]:%s\n", msg)

	   time.Sleep(time.Second)
	}
}
