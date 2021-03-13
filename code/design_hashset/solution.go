package design_hashset

type MyHashSet struct {
	inner []uint64
}

const tableSize = 1e6/64 + 1 // 15626

func Constructor() MyHashSet {
	return MyHashSet{
		inner: make([]uint64, tableSize),
	}
}

func (h *MyHashSet) Add(key int) {
	h.set(key, true)
}

func (h *MyHashSet) Remove(key int) {
	h.set(key, false)
}

func (h *MyHashSet) Contains(key int) bool {
	baseIndex, bitIndex := key/64, key%64
	mask := uint64(1) << bitIndex
	return h.inner[baseIndex]&mask != 0
}

func (h *MyHashSet) set(index int, value bool) {
	baseIndex, bitIndex := index/64, index%64
	mask := uint64(1) << bitIndex
	if value {
		h.inner[baseIndex] |= mask
	} else {
		h.inner[baseIndex] &= ^mask
	}
}
