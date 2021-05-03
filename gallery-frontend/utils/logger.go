package utils

import (
	"log"
	"os"
)

var InfoLog *log.Logger
var DebugLog *log.Logger
var ErrorLog *log.Logger

func init() {
	infoPrefix := "[INFO][gallery-frontend] "
	InfoLog = log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime)

	debugPrefix := "[DEBUG][gallery-frontend] "
	DebugLog = log.New(os.Stdout, debugPrefix, log.Ldate|log.Ltime|log.Llongfile)

	errorPrefix := "[ERROR][gallery-frontend] "
	ErrorLog = log.New(os.Stderr, errorPrefix, log.Ldate|log.Ltime|log.Llongfile)
}
