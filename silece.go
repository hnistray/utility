package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []string{"hello", "strict", "happy"}

	var wg sync.WaitGroup
	for _, str := range data {
		wg.Add(1)
		go func(str string) {
			fmt.Println(str)
			wg.Done()
		}(str)
	}

	wg.Wait()
}
