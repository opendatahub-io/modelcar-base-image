package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Hello, world! And thanks Jason for the idea, Daniele and Roland")
	checkIfEarlyReturn()
	fmt.Println("Creating symbolic link per ModelCar logic...")
	doTheThing()
	fmt.Println("Will now wait forever to serve as sidecar per ModelCar logic. (waiting forever)")
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	fmt.Println("Main terminated.")
}

func checkIfEarlyReturn() {
	fmt.Println("Invoked as:", strings.Join(os.Args, " "))
	shouldEarlyReturn := true
	for _, arg := range os.Args {
		if strings.Contains(arg, "sleep") {
			shouldEarlyReturn = false
		}
	}
	if shouldEarlyReturn {
		fmt.Println("Did not find any use of `sleep` in arguments, assuming InitContainer, terminating early with zero return code.")
		os.Exit(0)
	}
}

func doTheThing() {
	pid := os.Getpid()
	fromLink := filepath.Join("/proc", fmt.Sprintf("%d", pid), "root/models")
	toLink := "/mnt/models"

	if _, err := os.Lstat(toLink); err == nil { // see also https://github.com/kserve/kserve/pull/4274 for force recreation of symlink
		fmt.Printf("Link path %s already exists. Removing it.\n", toLink)
		err := os.Remove(toLink)
		if err != nil {
			fmt.Println("Error removing existing link:", err)
			return
		}
	}

	err := os.Symlink(fromLink, toLink)
	if err != nil {
		fmt.Println("Error creating symbolic link:", err)
		return
	}

	fmt.Println("Symbolic link created successfully from", fromLink, "to", toLink)
}
