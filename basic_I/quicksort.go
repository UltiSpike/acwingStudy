package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d\n", &n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &array[i])

	}

	q_sort(array, 0, n-1)

}

func q_sort1(array []int, l, r int) {
	if l >= r {
		return
	}

	pivot := array[rand.Intn(r-l+1)+l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; array[i] < pivot; i++ {
		}
		for j--; array[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}

	q_sort(array, l, j)
	q_sort(array, j+1, r)
	return
}

func partition(arr []int, l, r int) (int, int) {
	srand := rand.Intn(r-l+1) + l
	pivot := arr[srand]
	l--
	r++

	for {
		for l++; arr[l] < pivot; l++ {
		}
		for r--; arr[r] > pivot; r-- {
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	return pivot, r
}

func q_sort(array []int, l, r int) {
	if l < r {
		fmt.Println("start:", array[l:r+1])
		p1, p := partition(array, l, r)
		fmt.Println("pivot", p1, "end:", array[l:r+1])
		q_sort(array, l, p)
		q_sort(array, p+1, r)
	}
}

func parqSort(nums []int, l, r int, wg *sync.WaitGroup) {
	if l >= r {
		return
	}

	pivot := nums[rand.Intn(r-l+1)+l]
	i, j := l-1, r+1

	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		parqSort(nums, l, j, &wg1)
	}()
	go func() {
		defer wg1.Done()
		parqSort(nums, j+1, r, &wg1)
	}()
	wg1.Wait()
}
