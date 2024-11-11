package main

import (
	"fmt"
	"github.com/redis_golang/internal/network"
	"github.com/redis_golang/internal/storage"
)

func main() {
	db := storage.NewInMemoryDB()

	fmt.Println("Starting Redis server on port 8000...")

	server := network.NewTCPServer(":8000", db)
	server.Start()
}

// func main(){
// 	listener, err := net.Listen("tcp", ":8000")

// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	defer listener.Close()
// 	fmt.Println("server listening on port 8000...")

// 	for{
// 		conn, err := listener.Accept()
// 		if err != nil{
// 			fmt.Println("error 2")
// 			continue
// 		}

// 		fmt.Println("client Connected")

// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn){
// 	defer conn.Close()

// 	reader := bufio.NewReader(conn)
// 	for {
// 		message, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("error 3")
// 			return
// 		}
// 		fmt.Printf("message: %v",message)
// 		_, err = conn.Write([]byte("Hello\n"))
// 		if err != nil{
// 			fmt.Println("error sending message")
// 			return
// 		}
// 	}
// }
