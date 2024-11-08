package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s file", os.Args[0])
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()
	err = watcher.Add(filePath)
	if err != nil {
		log.Fatalf("Failed to add file to watcher: %v", err)
	}

	file.Seek(0, io.SeekEnd)
	reader = bufio.NewReader(file)

	wg.Add(1)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					for {
						line, err := reader.ReadString('\n')
						if err != nil {
							break
						}
						fmt.Print(line)
					}
				}
			case err := <-watcher.Errors:
				log.Print("Watcher error:", err)
			}
		}
	}()
	wg.Wait()
	fmt.Println("Main terminated.")
}
