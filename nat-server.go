package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:10086")
	if err != nil {
		fmt.Printf("udp resolve fail: err %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
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
	   buf := make([]byte, 4096)
	   n, raddr, err := conn.ReadFromUDP(buf)
	   if err != nil {
		  fmt.Printf("read fail: err %v", err)
	   }

	   buf = buf[:n]
	   fmt.Printf("[recv:%v]:%s\n", raddr, buf)
	}
}
