package tcp

import "net"

const (
	HOST = "127.0.0.1"
	PORT = "34070"
	TYPE = "tcp"
)

func StartServer() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("Message received."))
	conn.Close()
}
