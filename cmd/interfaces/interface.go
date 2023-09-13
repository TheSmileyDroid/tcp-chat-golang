package main

import (
	"net"
	"strconv"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			panic(err)
		}
		text := ""
		for _, addr := range addrs {
			text += " type: " + addr.Network() + " addr: " + addr.String()
		}
		println(strconv.Itoa(i)+":", inter.Name, inter.HardwareAddr.String(), text, "\n")
	}
}
