package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello, world"

	messages := []string{"Hello, universe", "Hello, cosmos", "Hello, everyone"}

	for _, i := range messages {
		wg.Add(1)
		go func(s string) {
			updateMessage(s)
			printMessage()
		}(i)
	}

	wg.Wait()

}
