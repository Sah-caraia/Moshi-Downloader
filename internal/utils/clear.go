package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	command := ""

	if runtime.GOOS == "windows" {
		command = "cls"
	} else {
		command = "clear"
	}
	
	cmdExec := exec.Command(command)
	cmdExec.Stdout = os.Stdout
	cmdExec.Run()
}
