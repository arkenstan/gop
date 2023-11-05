package main

import (
	"os"
	"os/exec"
)

func consoleClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
