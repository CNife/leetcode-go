package implement_trie_prefix_tree

type Trie struct {
	root node
}

type node struct {
	children [26]*node
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	n := &t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if n.children[idx] == nil {
			n.children[idx] = &node{}
		}
		n = n.children[idx]
	}
	n.isWord = true
}

func (t *Trie) Search(word string) bool {
	n := t.getNode(word)
	return n != nil && n.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.getNode(prefix) != nil
}

func (t *Trie) getNode(word string) *node {
	n := &t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if n.children[idx] == nil {
			return nil
		}
		n = n.children[idx]
	}
	return n
}
