package tcp

import "net"

func StartClient() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		panic(err)
	}

	println("Received: " + string(received) + "!")
	conn.Close()
}
