package ClearConsole

import (
	"os"
	"os/exec"
	"runtime"
)

func windows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func linux() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		windows()
	case "linux", "darwin":
		linux()
	}
}
