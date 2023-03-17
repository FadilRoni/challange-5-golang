package main

import (
	"fmt"
	"sync"
)


func main() {

	data1 := []string{"coba1", "coba2", "coba3"}

	data2 := []string{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go print(i, data1, data2, &mu, &wg)
	}
	wg.Wait()
}

func print(i int, data1 []string, data2 []string, mu *sync.Mutex, wg *sync.WaitGroup) {

	mu.Lock()	
	fmt.Println(data1, i)
	mu.Unlock()

	mu.Lock()
	fmt.Println(data2, i)
	mu.Unlock()

	wg.Done()
}