package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: %s addr:port", os.Args[0])
	}
	addr := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln("lookupzone Error is %s", err.Error())
	}
	fmt.Println("tcpAddr is", tcpAddr.String())
	fmt.Println("tcpAddr's zone is", tcpAddr.Zone)
}
