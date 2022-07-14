package prefix_and_suffix_search

type WordFilter struct {
	leftRoot, rightRoot trieNode
}

func Constructor(words []string) WordFilter {
	var wf WordFilter
	for i, word := range words {
		wf.leftRoot.addWord(word, i, true)
		wf.rightRoot.addWord(word, i, false)
	}
	return wf
}

func (wf *WordFilter) F(prefix, postfix string) int {
	prefixNode := wf.leftRoot.getNode(prefix, true)
	if prefixNode == nil {
		return -1
	}
	postfixNode := wf.rightRoot.getNode(postfix, false)
	if postfixNode == nil {
		return -1
	}
	return getCommonMaxOfSorted(prefixNode.idxs, postfixNode.idxs)
}

type trieNode struct {
	idxs  []int
	nexts [26]*trieNode
}

func (node *trieNode) addWord(word string, idx int, left2Right bool) {
	if left2Right {
		for i := 0; i < len(word); i++ {
			node = insertTrieNode(node, word[i], idx)
		}
	} else {
		for i := len(word) - 1; i >= 0; i-- {
			node = insertTrieNode(node, word[i], idx)
		}
	}
	node.idxs = append(node.idxs, idx)
}

func insertTrieNode(node *trieNode, b byte, idx int) *trieNode {
	node.idxs = append(node.idxs, idx)
	trieIndex := b - 'a'
	if node.nexts[trieIndex] == nil {
		node.nexts[trieIndex] = new(trieNode)
	}
	return node.nexts[trieIndex]
}

func (node *trieNode) getNode(word string, left2Right bool) *trieNode {
	if left2Right {
		for i := 0; i < len(word); i++ {
			trieIndex := word[i] - 'a'
			if node.nexts[trieIndex] != nil {
				node = node.nexts[trieIndex]
			} else {
				return nil
			}
		}
	} else {
		for i := len(word) - 1; i >= 0; i-- {
			trieIndex := word[i] - 'a'
			if node.nexts[trieIndex] != nil {
				node = node.nexts[trieIndex]
			} else {
				return nil
			}
		}
	}
	return node
}

func getCommonMaxOfSorted(lhs, rhs []int) int {
	i, j := len(lhs)-1, len(rhs)-1
	for i >= 0 && j >= 0 {
		if lhs[i] == rhs[j] {
			return lhs[i]
		} else if lhs[i] < rhs[j] {
			j--
		} else {
			i--
		}
	}
	return -1
}
