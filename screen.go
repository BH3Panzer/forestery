package main

import "os"
import "os/exec"

func Clear() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
