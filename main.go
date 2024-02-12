package main

import (
	"fmt"
	"os"
	"os/exec"
	//"github.com/diskfs/go-diskfs"
)

func ClearScr() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	ClearScr()
	fmt.Println("\033[1;32m>> \033[1;97mHello world!")
}
