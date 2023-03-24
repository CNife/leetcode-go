package stream_of_characters

type StreamChecker struct {
	trieRoot trieNode
	buffer   []byte
}

type trieNode struct {
	children [26]*trieNode
	isWord   bool
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{}
	for _, word := range words {
		sc.trieRoot.insertReversedWord(word)
	}
	return sc
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.buffer = append(sc.buffer, letter)
	return sc.trieRoot.searchPostfix(sc.buffer)
}

func (n *trieNode) insertReversedWord(word string) {
	for i := len(word) - 1; i >= 0; i-- {
		idx := word[i] - 'a'
		if n.children[idx] == nil {
			n.children[idx] = &trieNode{}
		}
		n = n.children[idx]
	}
	n.isWord = true
}

func (n *trieNode) searchPostfix(buffer []byte) bool {
	for i := len(buffer) - 1; i >= 0; i-- {
		idx := buffer[i] - 'a'
		if n.children[idx] == nil {
			return false
		}
		if n.children[idx].isWord {
			return true
		}
		n = n.children[idx]
	}
	return n.isWord
}
