package network

import (
	"fmt"
	"net"

	"github.com/redis_golang/internal/command"
	"github.com/redis_golang/internal/storage"
)

type TCPServer struct {
	address string
	db      *storage.InMemoryDB
}

func NewTCPServer(address string, db *storage.InMemoryDB) *TCPServer {
	return &TCPServer{address, db}
}

func (s *TCPServer) Start() {
	listener, err := net.Listen("tcp", s.address)

	if err != nil {
		fmt.Println("Error in starting tcp server")
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection")
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	command.HandleClient(conn, s.db)
}
