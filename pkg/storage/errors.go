package storage

import "fmt"

func PathError(msg string, rest ...interface{}) error {
	return fmt.Errorf(
		fmt.Sprintf("path error: %s", msg),
		rest...)
}
