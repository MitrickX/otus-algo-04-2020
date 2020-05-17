package queue

import (
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container"
)

// BinarySearch search in array by binary algorithm
// Function is generic so need know what how to compare with original value.
func BinarySearch(array container.Array, cmp func(x interface{}) int) (int, bool) {
	// length of array
	ln := array.Len()

	// left bound of search range
	left := 0

	// right bound of search range
	right := ln - 1

	// search while bounds not crossed
	for left <= right {
		// calc middle index
		middle := (left + right) / 2

		value := array.Get(middle)

		// hit - found what search
		if cmp(value) == 0 {
			return middle, true
		}

		// left and right bounds crossed, so it is miss case
		if left == right {
			// define place where need insert value
			if cmp(array.Get(left)) < 0 {
				return left + 1, false
			}

			return left, false
		}

		if cmp(value) < 0 {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	// empty array case
	return 0, false
}

// BinarySearchInt for array of integers.
func BinarySearchInt(array container.Array, val int) (int, bool) {
	return BinarySearch(array, func(x interface{}) int {
		return x.(int) - val
	})
}
