package cli

import (
	"os"
	"os/exec"
)

func Run(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		Cry("error starting: %v", err)
	}
	return cmd
}

func CaptureOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		Cry("error running: %v", err)
	}
	return string(out)
}
