package main

import (
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, inter := range interfaces {
		log.Printf("%d: %s %s", i, inter.Name, inter.HardwareAddr)
	}
}
