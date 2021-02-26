package intersection_of_two_arrays

import "sort"

func Intersection(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var result []int
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		lhs, rhs := arr1[i], arr2[j]
		switch {
		case lhs < rhs:
			i++
		case lhs > rhs:
			j++
		default:
			if len(result) == 0 || result[len(result)-1] != lhs {
				result = append(result, lhs)
			}
			i++
			j++
		}
	}
	return result
}
