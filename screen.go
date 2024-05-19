package main

import "os"
import "os/exec"
import "runtime"

func Clear() {
	if runtime.GOOS == "windows" {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	} else if runtime.GOOS == "linux" {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	}
}
