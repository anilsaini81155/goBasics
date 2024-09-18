/*Program to print sum of square of the numbers:
input:
arr = [1,2,3,4,5,6,7,8,9]
n = 4
n is the number of goroutines to be created for generating squares in random order of arr.
create a consumer goroutine to receive this squares and sum it up.
main function should recieve this final sum and print it.
use channels where ever necessary.
*/

package main

import (
	"fmt"
	"sync"
)

// Producer function to calculate squares and send them to the channel
func producer(nums []int, squareCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		square := num * num
		squareCh <- square
	}
}

// Consumer function to sum up squares from the channel
func consumer(squareCh chan int, sumCh chan int) {
	sum := 0
	for square := range squareCh {
		sum += square
	}
	// Send the final sum to the sum channel
	sumCh <- sum
	// Close sumCh to signal completion
	close(sumCh)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 4 // Number of producer goroutines
	squareCh := make(chan int)
	sumCh := make(chan int)
	var wg sync.WaitGroup

	// Divide the array into chunks for each producer goroutine
	chunkSize := (len(arr) + n - 1) / n

	// Start producer goroutines
	for i := 0; i < n; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		wg.Add(1)
		go producer(arr[start:end], squareCh, &wg)
	}

	// Start consumer goroutine
	go consumer(squareCh, sumCh)

	// Wait for all producer goroutines to finish
	wg.Wait()
	// Close the square channel to signal the consumer to finish
	close(squareCh)

	// Receive the final sum and print it
	finalSum := <-sumCh
	fmt.Println("Sum of squares:", finalSum)
}
