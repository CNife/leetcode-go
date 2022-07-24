package distance_between_bus_stops

func DistanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	d1 := 0
	for i := start; i < destination; i++ {
		d1 += distance[i]
	}
	d2 := 0
	for i := 0; i < start; i++ {
		d2 += distance[i]
	}
	for i := destination; i < len(distance); i++ {
		d2 += distance[i]
	}

	if d1 < d2 {
		return d1
	}
	return d2
}
