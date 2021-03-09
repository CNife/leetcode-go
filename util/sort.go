package util

import "sort"

func SortInts(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func SortStrings(strs []string) []string {
	sort.Strings(strs)
	return strs
}

func SortIntSlices(slices [][]int) [][]int {
	return SortIntSlicesDeep(slices, func(from []int) []int {
		return from
	})
}

func SortIntSlicesDeep(
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

func SortStringSlices(slices [][]string) [][]string {
	return SortStringSlicesDeep(slices, func(from []string) []string {
		return from
	})
}

func SortStringSlicesDeep(
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
