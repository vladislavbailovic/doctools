package cli

import "os"

func HasFlag(flag string) bool {
	for _, a := range os.Args {
		if a == flag {
			return true
		}
	}
	return false
}

func HasSubcommand() bool {
	return len(os.Args) > 1
}

func HasSubcommandArgs() bool {
	return len(os.Args) > 2
}

func Subcommand() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ""
}

func SubcommandArgs() []string {
	return ArgsFrom(2)
}

func ArgsFrom(where int) []string {
	if len(os.Args) < where {
		return []string{}
	}
	return os.Args[where:]
}
