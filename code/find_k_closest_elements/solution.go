package find_k_closest_elements

func FindClosestElements(arr []int, k, x int) []int {
	nearest := findNearest(arr, x)
	i, j := nearest-1, nearest+1
	for j-i-1 < k {
		if i >= 0 && j < len(arr) {
			if distance(x, arr[i]) <= distance(x, arr[j]) {
				i--
			} else {
				j++
			}
		} else if i >= 0 {
			i--
		} else {
			j++
		}
	}

	result := make([]int, j-i-1)
	copy(result, arr[i+1:j])
	return result
}

func findNearest(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low < high {
		middle := (low + high + 1) / 2
		if arr[middle] <= x {
			low = middle
		} else {
			high = middle - 1
		}
	}
	if high+1 < len(arr) && distance(x, arr[high+1]) < distance(x, arr[high]) {
		return high + 1
	}
	return high
}

func distance(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
