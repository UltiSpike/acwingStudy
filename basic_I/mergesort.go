package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d \n", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}

	//mergesort(nums, 0, n-1)
	mergesort_(nums, n)
	for _, v := range nums {
		fmt.Printf("%d", v)
	}
}

func mergesort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort(nums, l, mid)
	mergesort(nums, mid+1, r)
	merge(nums, l, r)
}

func merge(nums []int, l, r int) {

	mid := (l + r) / 2
	temp := make([]int, r-l+1)
	i := l
	j := mid + 1
	p := 0
	for i <= mid && j <= r {

		if nums[i] <= nums[j] {
			temp[p] = nums[i]
			i++
		} else {
			temp[p] = nums[j]
			j++
		}
		p++

	}
	for i <= mid {
		temp[p] = nums[i]
		i++
		p++
	}
	for j <= r {
		temp[p] = nums[j]
		j++
		p++
	}
	p = 0

	for i = l; i <= r; i++ {
		nums[i] = temp[p]
		p++
	}
}

// 非递归式归并排序(错误的）
func mergesort_(nums []int, n int) {
	t := 1
	for t < n {
		for low := 0; low < n; low += 2 * t {
			mid := min(low+t, n-1)
			high := min(low+2*t-1, n-1)
			tmp := merge_(nums[low:mid+1], nums[mid:high+1])
			for i := low; i < high+1; i++ {
				nums[i] = tmp[i-low]

			}
			fmt.Println(nums, tmp)
		}
		t *= 2
	}
}

// 非递归式merge 需要指明切分点
func merge_(numsA, numsB []int) []int {
	m := len(numsA) - 1
	n := len(numsB) - 1
	var ans []int
	i, j := 0, 0
	for i <= m && j <= n {
		if numsA[i] <= numsB[j] {
			ans = append(ans, numsA[i])
			i++
		} else {
			ans = append(ans, numsB[j])
			j++
		}
	}
	for i <= m {
		ans = append(ans, numsA[i])
		i++
	}
	for j <= n {
		ans = append(ans, numsB[j])
		j++
	}

	return ans
}
