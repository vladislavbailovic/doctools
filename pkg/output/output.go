package output

import (
	"fmt"
	"os"
)

func Cry(msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stderr,
		fmt.Sprintf("[ERROR] %s\n", msg), rest...)
}

func Nit(msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stderr,
		fmt.Sprintf("[DEBUG] %s\n", msg), rest...)
}

func Say(msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", rest...)
}
