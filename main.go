package main

import (
	"fmt"
	// "github.com/diskfs/go-diskfs"
	"os"
	"os/exec"
	"os/user"
)

func ClearScr() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func isRoot() bool {
    currentUser, err := user.Current()
    if err != nil {
        fmt.Println("\033[1;31m==> \033[97mUnable to get user.")
    }
    return currentUser.Username == "root"
}

func main() {
	ClearScr()	
	var isrootuser bool = isRoot()

	if !isrootuser {
		fmt.Println("\033[1;31m==> \033[97mNot running as root. Please try again as root.")
		os.Exit(1)
	}

	fmt.Println("\033[1;32m==> \033[97mHello world!")
	
}
