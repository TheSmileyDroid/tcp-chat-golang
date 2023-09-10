package tcp

import (
	"bufio"
	"net"
	"os"
)

func StartClient() {
	HOST := os.Getenv("HOST")

	if HOST == "" {
		HOST = "localhost"
	}

	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)

		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			os.Stdout.WriteString(message)
		}
	}()

	os.Stdout.WriteString("Digite seu usuário: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		conn.Write([]byte(message))
	}
}
