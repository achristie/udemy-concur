package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"tau",
		"zeta",
		"epsilon",
		"theta",
	}

	for i, x := range words {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printSomething(fmt.Sprintf("%d: %s", i, x))

		}()
	}

	wg.Wait()

}
