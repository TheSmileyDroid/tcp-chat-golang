package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	HOST = "127.0.0.1"
	PORT = "34070"
	TYPE = "tcp"
)

type Client struct {
	Id   int
	Name string
}

func StartServer() {

	clientCount := 0

	clients := make(map[net.Conn]Client)

	connecting := make(chan net.Conn)

	dead := make(chan net.Conn)

	messages := make(chan string)

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				panic(err)
			}

			connecting <- conn
		}
	}()

	for {
		select {
		case conn := <-connecting:
			go func() {
				username, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Printf("Error reading username from client: %v", err)
					return
				}

				clients[conn] = Client{Id: clientCount, Name: strings.TrimSpace(username)}
				clientCount++

				log.Printf("%s conectou-se", clients[conn].Name)
				message := fmt.Sprintf("%s conectou-se\n", clients[conn].Name)
				sendToAll(clients, message, dead)

				go listenMessages(conn, messages, clients[conn], dead)
			}()

		case conn := <-dead:
			log.Printf("%s desconectou-se", clients[conn].Name)
			message := fmt.Sprintf("%s desconectou-se\n", clients[conn].Name)
			sendToAll(clients, message, dead)
			delete(clients, conn)

		case message := <-messages:
			sendToAll(clients, message, dead)

			log.Printf("%s", message)
		}
	}
}

func sendTo(conn net.Conn, message string, dead chan net.Conn) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		dead <- conn
		return
	}
}

func sendToAll(clients map[net.Conn]Client, message string, dead chan net.Conn) {
	for conn := range clients {
		go sendTo(conn, message, dead)
	}
}

func listenMessages(conn net.Conn, messages chan string, client Client, dead chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		messages <- fmt.Sprintf("%s: %s", client.Name, message)
	}

	dead <- conn
}
