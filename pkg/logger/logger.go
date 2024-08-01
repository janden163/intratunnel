package logger

import (
	"log"
	"os"
)

func Init(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
}

func Info(format string, v ...interface{}) {
	log.Printf("INFO: "+format, v...)
}

func Error(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}
