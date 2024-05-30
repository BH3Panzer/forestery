package main

import "os"
import "os/exec"
import "runtime"
import "fmt"

func Clear() {
	if runtime.GOOS == "windows" {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	} else if runtime.GOOS == "linux" {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	} else {
		fmt.Println("Platform unsupported :(")
	}
}

func GetAnsiColor(color string) string {
	if color == "black" {
		return "\033[30m"
	} else if color == "red" {
		return "\033[31m"
	} else if color == "green" {
		return "\033[32m"
	} else if color == "yellow" {
		return "\033[33m"
	} else if color == "blue" {
		return "\033[34m"
	} else if color == "magenta" {
		return "\033[35m"
	} else if color == "cyan" {
		return "\033[36m"
	} else if color == "white" {
		return "\033[37m"
	} else {
		return "\033[0m"
	}
}
