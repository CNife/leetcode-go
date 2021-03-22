package number_of_1_bits

func HammingWeight(num uint32) int {
	const (
		m0 uint32 = 0x55555555
		m1 uint32 = 0x33333333
		m2 uint32 = 0x0f0f0f0f
		m3 uint32 = 0x00ff00ff
		m4 uint32 = 0x0000ffff
	)

	num = (num & m0) + ((num >> 1) & m0)
	num = (num & m1) + ((num >> 2) & m1)
	num = (num & m2) + ((num >> 4) & m2)
	num = (num & m3) + ((num >> 8) & m3)
	num = (num & m4) + ((num >> 16) & m4)
	return int(num)
}
