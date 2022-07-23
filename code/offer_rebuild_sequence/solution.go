package offer_rebuild_sequence

func SequenceReconstruction(nums []int, sequences [][]int) bool {
	routes := make(map[int]map[int]struct{})
	for _, sequence := range sequences {
		for i := 1; i < len(sequence); i++ {
			prev, curr := sequence[i-1], sequence[i]
			if route, exists := routes[prev]; exists {
				route[curr] = struct{}{}
			} else {
				routes[prev] = map[int]struct{}{curr: {}}
			}
		}
	}
	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]
		if route, exists := routes[prev]; exists {
			if _, nextExists := route[curr]; nextExists {
				continue
			}
		}
		return false
	}
	return true
}
