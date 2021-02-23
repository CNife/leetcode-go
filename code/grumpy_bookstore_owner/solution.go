package grumpy_bookstore_owner

func MaxSatisfied(customers, grumpy []int, x int) int {
	annoyedCustomers := 0
	for i, customer := range customers[:x] {
		if grumpy[i] == 1 {
			annoyedCustomers += customer
		}
	}
	maxAnnoyedCustomers := annoyedCustomers

	for i := 0; i < len(customers)-x; i++ {
		if grumpy[i] == 1 {
			annoyedCustomers -= customers[i]
		}

		j := i + x
		if grumpy[j] == 1 {
			annoyedCustomers += customers[j]
		}

		maxAnnoyedCustomers = max(maxAnnoyedCustomers, annoyedCustomers)
	}

	safeCustomers := 0
	for i, customer := range customers {
		if grumpy[i] == 0 {
			safeCustomers += customer
		}
	}

	return maxAnnoyedCustomers + safeCustomers
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
