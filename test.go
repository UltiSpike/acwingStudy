package main

import "fmt"

func main() {
	table := map[int]int{}

	nums := []int{1, 2, 3}

	for i, v := range nums {
		table[v] = i
	}
	for i, v := range nums {
		table[v] += i
	}

	fmt.Println(table)
}
