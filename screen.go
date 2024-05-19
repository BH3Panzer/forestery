package main

import "os"
import "os/exec"
import "runtime"

func getRightCommand() string {
	if runtime.GOOS == "windows" {
		return "cls"
	} else if runtime.GOOS == "linux" {
		return "clear"
	} else {
		return "clear"
	}
}

func Clear() {
	com := getRightCommand()
	clear := exec.Command(com)
	clear.Stdout = os.Stdout
	clear.Run()
}
