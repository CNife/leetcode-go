package sort

import "sort"

func Ints(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func Strings(strs []string) []string {
	sort.Strings(strs)
	return strs
}

func IntSlices(slices [][]int) [][]int {
	return IntSlicesDeep(slices, func(from []int) []int {
		return from
	})
}

func IntSlicesDeep(
	slices [][]int,
	mapFunc func(from []int) []int,
) [][]int {
	for i, slice := range slices {
		slices[i] = mapFunc(slice)
	}

	sort.Slice(slices, func(i, j int) bool {
		lhs, rhs := slices[i], slices[j]
		if len(lhs) != len(rhs) {
			return len(lhs) < len(rhs)
		}
		for k := 0; k < len(lhs); k++ {
			if lhs[k] != rhs[k] {
				return lhs[k] < rhs[k]
			}
		}
		return false
	})

	return slices
}

func StringSlices(slices [][]string) [][]string {
	return StringSlicesDeep(slices, func(from []string) []string {
		return from
	})
}

func StringSlicesDeep(
	slices [][]string,
	mapFunc func(from []string) []string,
) [][]string {
	for i, slice := range slices {
		slices[i] = mapFunc(slice)
	}

	sort.Slice(slices, func(i, j int) bool {
		lhs, rhs := slices[i], slices[j]
		if len(lhs) != len(rhs) {
			return len(lhs) < len(rhs)
		}
		for k := 0; k < len(lhs); k++ {
			if lhs[k] != rhs[k] {
				return lhs[k] < rhs[k]
			}
		}
		return false
	})

	return slices
}
