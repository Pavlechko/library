package utils

import (
	"log"
	"os"
	"path/filepath"
)

const chmod = 0666

var InfoLogger *log.Logger
var ErrorLogger *log.Logger

func init() {
	pathSeck := filepath.FromSlash("logs/books-log.log")
	myLog, err := os.OpenFile(pathSeck, os.O_RDWR|os.O_CREATE|os.O_APPEND, chmod)
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}
	InfoLogger = log.New(myLog, "[Info]:\t", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	ErrorLogger = log.New(myLog, "[Error]:\t", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
}
