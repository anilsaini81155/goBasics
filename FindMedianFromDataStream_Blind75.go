/*

The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
 

Example 1:

Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
 

Constraints:

-105 <= num <= 105
There will be at least one element in the data structure before calling findMedian.
At most 5 * 104 calls will be made to addNum and findMedian.

//Implement in golang


///Below is the given structure

type MedianFinder struct {
    
}


func Constructor() MedianFinder {
    
}


func (this *MedianFinder) AddNum(num int)  {
    
}


func (this *MedianFinder) FindMedian() float64 {
    
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

//
// Actual Code Starts here


package main

import "fmt"

type MedianFinder struct {
	lower []int // max-heap (manually managed)
	upper []int // min-heap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	// Insert into appropriate heap
	if len(this.lower) == 0 || num <= this.lower[0] {
		this.lower = insertMax(this.lower, num)
	} else {
		this.upper = insertMin(this.upper, num)
	}

	// Balance the heaps
	if len(this.lower) > len(this.upper)+1 {
		val, newLower := removeMax(this.lower)
		this.lower = newLower
		this.upper = insertMin(this.upper, val)
	} else if len(this.upper) > len(this.lower) {
		val, newUpper := removeMin(this.upper)
		this.upper = newUpper
		this.lower = insertMax(this.lower, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.lower) > len(this.upper) {
		return float64(this.lower[0])
	}
	return float64(this.lower[0]+this.upper[0]) / 2.0
}

// Min-heap
func insertMin(h []int, val int) []int {
	h = append(h, val)
	i := len(h) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h[i] < h[p] {
			h[i], h[p] = h[p], h[i]
			i = p
		} else {
			break
		}
	}
	return h
}

func removeMin(h []int) (int, []int) {
	n := len(h)
	h[0], h[n-1] = h[n-1], h[0]
	val := h[n-1]
	h = h[:n-1]
	i := 0
	for {
		smallest := i
		l, r := 2*i+1, 2*i+2
		if l < len(h) && h[l] < h[smallest] {
			smallest = l
		}
		if r < len(h) && h[r] < h[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h[i], h[smallest] = h[smallest], h[i]
		i = smallest
	}
	return val, h
}

// Max-heap
func insertMax(h []int, val int) []int {
	h = append(h, val)
	i := len(h) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h[i] > h[p] {
			h[i], h[p] = h[p], h[i]
			i = p
		} else {
			break
		}
	}
	return h
}

func removeMax(h []int) (int, []int) {
	n := len(h)
	h[0], h[n-1] = h[n-1], h[0]
	val := h[n-1]
	h = h[:n-1]
	i := 0
	for {
		largest := i
		l, r := 2*i+1, 2*i+2
		if l < len(h) && h[l] > h[largest] {
			largest = l
		}
		if r < len(h) && h[r] > h[largest] {
			largest = r
		}
		if largest == i {
			break
		}
		h[i], h[largest] = h[largest], h[i]
		i = largest
	}
	return val, h
}

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian()) // 1.5
	m.AddNum(3)
	fmt.Println(m.FindMedian()) // 2.0
}

