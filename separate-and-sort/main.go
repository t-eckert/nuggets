package main

import (
	"fmt"
	"sort"
)

func separateAndSort(input []int) ([]int, []int) {
	even, odd := make([]int, 0), make([]int, 0)

	for _, v := range input {
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	sort.Ints(even)
	sort.Ints(odd)

	return even, odd
}

func test(input, even, odd []int) bool {
	actualEven, actualOdd := separateAndSort(input)

	return areEqual(even, actualEven) && areEqual(odd, actualOdd)
}

func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func main() {
	cases := []struct {
		input []int
		even  []int
		odd   []int
	}{
		{[]int{4, 3, 2, 1, 5, 7, 6, 9}, []int{2, 4, 6}, []int{1, 3, 5, 7, 9}},
		{[]int{1, 1, 1, 1}, []int{}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}, []int{}},
	}

	for _, c := range cases {
		if !test(c.input, c.even, c.odd) {
			fmt.Println(c.input, c.even, c.odd)
			panic("Failed!")
		}
	}
}
