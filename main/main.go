package main

import (
	"fmt"
	// "github.com/diskfs/go-diskfs"
	"os"
	"os/exec"
	// "encoding/json"
)

func ClearScr() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	ClearScr()
	fmt.Println("\033[1;32m==> \033[97mHello world.")
	os.Exit(0);
}
