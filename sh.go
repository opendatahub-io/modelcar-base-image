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
	fmt.Println("Invoked with and will check later is consistent with KServe deployment", strings.Join(os.Args[1:], " ")) // TODO will need to check the args here.
	doTheThing()
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	fmt.Println("Main terminated.")
}

func doTheThing() {
	pid := os.Getpid()
	fromLink := filepath.Join("/proc", fmt.Sprintf("%d", pid), "root/models")
	toLink := "/mnt/models"

	if _, err := os.Lstat(toLink); err == nil {
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
