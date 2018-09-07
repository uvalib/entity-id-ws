package logger

import (
	"log"
)

//
// Log -- our logger...
//
func Log(msg string) {
	log.Printf("%s", msg)
}

//
// end of file
//
