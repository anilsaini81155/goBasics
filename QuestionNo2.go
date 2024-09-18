/*Find the duplicate number in a sorted array that has n elements. The array will contain all
the numbers ranging from 1 to n-1.
E.g
{1,2,2,3,4,5, 6} = 2
{1,2,3,4,4,5,6} = 4
*/

package main

import (
	"fmt"
)

func findDuplicate(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			fmt.Println(i)
			return arr[i]
		}
	}
	return -1
}

func main() {

	arr1 := []int{1, 2, 2, 3, 4, 5, 6}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Dupliacte No Found Result==>", findDuplicate(arr1))
	fmt.Println("Dupliacte No Found Result==>", findDuplicate(arr2))

}
