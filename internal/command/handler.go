package command

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/redis_golang/internal/storage"
)

func HandleClient(conn net.Conn, db *storage.InMemoryDB) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		message = strings.TrimSpace(message)
		if len(message) == 0 {
			continue
		}

		response := ExecuteCommand(message, db)
		conn.Write([]byte(response+"\n"))
	}
}
