package capacity_to_ship_packages_within_d_days

import "sort"

func ShipWithinDays(weights []int, day int) int {
	min, max := 0, 0
	for _, weight := range weights {
		max += weight
		if weight > min {
			min = weight
		}
	}

	result := sort.Search(max-min+1, func(i int) bool {
		sum, counter := 0, 1
		for _, weight := range weights {
			sum += weight
			if sum > i+min {
				sum = weight
				counter++
				if counter > day {
					return false
				}
			}
		}
		return true
	})
	return result + min
}
