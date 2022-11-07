package cli

import "os"

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
