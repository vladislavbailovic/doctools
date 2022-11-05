package config

import "fmt"

func LoadError(msg string, rest ...interface{}) error {
	return fmt.Errorf(
		fmt.Sprintf("unable to load config: %s", msg),
		rest...)
}

func InitError(msg string, rest ...interface{}) error {
	return fmt.Errorf(
		fmt.Sprintf("unable to initialize config: %s", msg),
		rest...)
}
