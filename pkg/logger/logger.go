package logger;

import (
	"log"
	"os"
)

var(
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func info(msg string){
	logger.Println(msg)
}