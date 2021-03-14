package design_hashmap

type MyHashMap struct {
	table []int
}

func Constructor() MyHashMap {
	table := make([]int, 1e6+1)
	for i := range table {
		table[i] = -1
	}
	return MyHashMap{table}
}

func (h *MyHashMap) Put(key, value int) {
	h.table[key] = value
}

func (h *MyHashMap) Get(key int) int {
	return h.table[key]
}

func (h *MyHashMap) Remove(key int) {
	h.Put(key, -1)
}

//
// type MyHashMap struct {
// 	table []*hashMapEntry
// 	size  int
// }
//
// type hashMapEntry struct {
// 	key, value int
// 	next       *hashMapEntry
// }
//
// const loadFactor = 0.75
//
// func Constructor() MyHashMap {
// 	return MyHashMap{
// 		table: make([]*hashMapEntry, 8),
// 	}
// }
//
// func (h *MyHashMap) Put(key, value int) {
// 	if float64(h.size)/float64(len(h.table)) > loadFactor {
// 		h.rehash()
// 	}
//
// 	putEntry(h.table, key, value)
// 	h.size++
// }
//
// func (h *MyHashMap) Get(key int) int {
// 	for ptr := *bucketPtr(h.table, key); ptr != nil; ptr = ptr.next {
// 		if ptr.key == key {
// 			return ptr.value
// 		}
// 	}
// 	return -1
// }
//
// func (h *MyHashMap) Remove(key int) {
// 	bucket := bucketPtr(h.table, key)
// 	if *bucket == nil {
// 		return
// 	}
//
// 	cur := *bucket
// 	if cur.key == key {
// 		*bucket = nil
// 		h.size--
// 		return
// 	}
//
// 	pre := cur
// 	cur = cur.next
// 	for cur != nil {
// 		if cur.key == key {
// 			pre.next = cur.next
// 			h.size--
// 			return
// 		}
// 		pre, cur = cur, cur.next
// 	}
// }
//
// func (h *MyHashMap) rehash() {
// 	newTable := make([]*hashMapEntry, 2*len(h.table))
// 	for _, entry := range h.table {
// 		for ; entry != nil; entry = entry.next {
// 			putEntry(newTable, entry.key, entry.value)
// 		}
// 	}
// 	h.table = newTable
// }
//
// func bucketPtr(table []*hashMapEntry, key int) **hashMapEntry {
// 	index := key % len(table)
// 	return &table[index]
// }
//
// func putEntry(table []*hashMapEntry, key, value int) {
// 	bucket := bucketPtr(table, key)
// 	if *bucket == nil {
// 		*bucket = &hashMapEntry{key, value, nil}
// 		return
// 	}
//
// 	pre, cur := (*hashMapEntry)(nil), *bucket
// 	for cur != nil {
// 		if cur.key == key {
// 			cur.value = value
// 			return
// 		}
// 		pre, cur = cur, cur.next
// 	}
// 	pre.next = &hashMapEntry{key, value, nil}
// }
