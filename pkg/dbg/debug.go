package dbg

import (
	"fmt"
	"log"
)

func Error(msg string, rest ...interface{}) {
	dbg(fmt.Sprintf("ERROR: %s", msg), rest...)
}

func Debug(msg string, rest ...interface{}) {
	dbg(fmt.Sprintf("DEBUG: %s", msg), rest...)
}

func dbg(msg string, rest ...interface{}) {
	log.Println(fmt.Sprintf(msg, rest...))
}

func PathError(msg string, rest ...interface{}) error {
	return fmt.Errorf(
		fmt.Sprintf("path error: %s", msg),
		rest...)
}
