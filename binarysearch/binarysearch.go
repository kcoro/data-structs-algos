package binarysearch

/*
	binarysearch finds a value in a sorted list in O(logN).
	Binary search looks at the middle index of a sorted linear data structure to search
	for a value.  If the value is not found, binary search checks if the value is higher or lower
	than the value at the midpoint. Each iteration the value being searched is not found, binary search
	divides the remaining number of elements to search by 1 / 2.
*/
func binarysearch(find int, list []int) int {
	low := 0
	high := len(list) - 1
	mid := (low + high) / 2

	for low <= high {
		if list[mid] == find {
			return mid
		} else if find > list[mid] {
			low = mid + 1
			mid = low + high/2
		} else if find < list[mid] {
			high = mid - 1
			mid = low + high/2
		}
	}

	// -1 signals no valid value found
	return -1
}
