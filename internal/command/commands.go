package command

import (
	"fmt"
	"strings"

	"github.com/redis_golang/internal/storage"
)

func ExecuteCommand(message string, db *storage.InMemoryDB) string {
	parts := strings.SplitN(message, " ", 3)
	command := strings.ToUpper(parts[0])

	switch command {
	case "GET":
		if len(parts) < 2 {
			return "ERR wrong number of arguments for 'get' command"
		}
		key := parts[1]
		value, exists := db.Get(key)
		if !exists {
			return "(nil)"
		}
		return value
	case "SET":
		if len(parts) < 3 {
			return "ERR wrong number of arguments for 'set' command"
		}
		key := parts[1]
		value := parts[2]
		db.Set(key, value)
		return "OK"
	default:
		return fmt.Sprintf("ERR Unknown command '%s'", command)
	}
}
