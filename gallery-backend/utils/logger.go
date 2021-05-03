package utils

import (
	"log"
	"os"
)

var InfoLog *log.Logger
var DebugLog *log.Logger
var ErrorLog *log.Logger

func init() {
	infoPrefix := "[INFO][gallery-backend] "
	InfoLog = log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime)

	debugPrefix := "[DEBUG][gallery-backend] "
	DebugLog = log.New(os.Stdout, debugPrefix, log.Ldate|log.Ltime|log.Llongfile)

	errorPrefix := "[ERROR][gallery-backend] "
	ErrorLog = log.New(os.Stderr, errorPrefix, log.Ldate|log.Ltime|log.Llongfile)
}
