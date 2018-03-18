package main

import (
	"fmt"
	"github.com/go-fsnotify/fsnotify"
	"log"
	"syscall"
)

func main() {
	buff := make([]byte, 64)
	inotefd, err := syscall.InotifyInit()
	if err != nil {
		fmt.Println(err)
	}
	_, err = syscall.InotifyAddWatch(inotefd, "/home/nokia/NAS/testfile123.mp4", syscall.IN_CLOSE_WRITE)
	if err != nil {
		fmt.Println(err)
	}

	for {
		n, err := syscall.Read(inotefd, buff)
		if err != nil {
			fmt.Println(err)
			return
		}

		if n < 0 {
			fmt.Println("Read Error")
			return
		}

		fmt.Printf("Buffer: %v\n", buff)
		//can't cast []buff into InotifyEvent struct
		fmt.Printf("Cookie: %v\n", buff[0:4])
		fmt.Printf("Len: %v\n", buff[4:8])
		fmt.Printf("Mask: %v\n", buff[8:12])
		fmt.Printf("Name: %v\n", buff[12:13])
		fmt.Printf("Wd: %v\n", buff[13:17])
	}
}

func fsnotifydemo(){
	// Based on https://github.com/fsnotify/fsnotify
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				if event.Op & fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				fmt.Printf("EVENT! %#v\n", event)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("/Users/skdomino/Desktop/test.html"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
