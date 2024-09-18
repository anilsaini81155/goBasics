//Print numbers from 1 to 10 using 2 goroutines. One printing even and the other printing odd numbers.
package main

import (
	"fmt"
	"sync"
)

func oddPrint(chOdd, chEven chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-chOdd
		fmt.Println(i)
		chEven <- true
	}
}

func evenPrint(chOdd, chEven chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-chEven
		fmt.Println(i)
		chOdd <- true
	}
}

func main() {

	chOdd := make(chan bool)
	chEven := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go oddPrint(chOdd, chEven, &wg)
	go evenPrint(chOdd, chEven, &wg)

	chEven <- true

	wg.Wait()
	close(chOdd)
	close(chEven)

}
