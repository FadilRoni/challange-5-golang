package main

import (
	"fmt"
	"sync"
)


func main() {

	data1 := []string{"coba1", "coba2", "coba3"}

	data2 := []string{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go coba(i, data1, &wg)
		wg.Add(1)
		go bisa(i, data2, &wg)
	}
	wg.Wait()
}

func coba(i int, data []string, wg *sync.WaitGroup) {
		fmt.Println(data, i)
		wg.Done()
}

func bisa(i int, data []string, wg *sync.WaitGroup) {
	fmt.Println(data, i)
	wg.Done()
}