package main

import (
	"fmt"
	// "github.com/diskfs/go-diskfs"
	"os"
	"os/exec"
	"os/user"
	"regexp"
)

func ClearScr() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("\033[1;31m==> \033[97mUnable to get user mode. Aborting.")
		os.Exit(2)
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

	lscmd := exec.Command("ls", "/dev")
	lsout, lserr := lscmd.CombinedOutput()

	if lserr != nil {
		fmt.Println("\033[1;31m==> \033[97mCommand threw error. Cannot continue.")
		os.Exit(3)
	}

	lsoutput := string(lsout)
	lscmd.Run()

	pattern := `sd[a-z][0-9]`

	re := regexp.MustCompile(pattern)

	drives := re.FindStringSubmatch(lsoutput)

	for _, drive := range drives {
		// Use the regular expression to check for a match
		if re.MatchString(drive) {
			fmt.Printf("%s\n", drive)
		} else {
			fmt.Printf("%s\n", drive)
		}
	}
}
