package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}
	sort.Slice(numStrs, func(i, j int) bool {
		lhs, rhs := numStrs[i], numStrs[j]
		return lhs+rhs > rhs+lhs
	})

	result := strings.Join(numStrs, "")
	if result[0] == '0' {
		return "0"
	}
	return result
}

//
// import (
// 	"container/heap"
// 	"strconv"
// 	"strings"
// )
//
// func LargestNumber(nums []int) string {
// 	numHeap := strs(make([]string, 0, len(nums)))
// 	for _, num := range nums {
// 		heap.Push(&numHeap, strconv.Itoa(num))
// 	}
//
// 	var sb strings.Builder
// 	for numHeap.Len() > 0 {
// 		sb.WriteString(numHeap.Pop().(string))
// 	}
//
// 	result := sb.String()
// 	if result[0] == '0' {
// 		return "0"
// 	}
// 	return result
// }
//
// type strs []string
//
// func (ss strs) Len() int {
// 	return len(ss)
// }
//
// func (ss strs) Less(i, j int) bool {
// 	lhs, rhs := ss[i], ss[j]
// 	m, n := len(lhs), len(rhs)
// 	for i, j := 0, 0; i < m+n; i, j = i+1, j+1 {
// 		var left, right byte
// 		if i < m {
// 			left = lhs[i]
// 		} else {
// 			left = rhs[i-m]
// 		}
// 		if j < n {
// 			right = rhs[j]
// 		} else {
// 			right = lhs[j-n]
// 		}
// 		if left != right {
// 			return left < right
// 		}
// 	}
// 	return false
// }
//
// func (ss strs) Swap(i, j int) {
// 	ss[i], ss[j] = ss[j], ss[i]
// }
//
// func (ss *strs) Push(x interface{}) {
// 	*ss = append(*ss, x.(string))
// }
//
// func (ss *strs) Pop() interface{} {
// 	old := *ss
// 	n := len(old)
// 	x := old[n-1]
// 	*ss = old[:n-1]
// 	return x
// }
