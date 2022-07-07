package replace_words

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	trieRoot := &trieNode{}
	for _, word := range dictionary {
		trieRoot.insert(word)
	}

	sentenceWords := strings.Split(sentence, " ")
	for i, word := range sentenceWords {
		prefix := trieRoot.findWordPrefix(word)
		if len(prefix) > 0 {
			sentenceWords[i] = prefix
		}
	}
	return strings.Join(sentenceWords, " ")
}

type trieNode struct {
	word      string
	nextNodes [26]*trieNode
}

func (node *trieNode) findWordPrefix(word string) string {
	ptr := node
	for i := 0; i < len(word); i++ {
		if ptr == nil {
			return ""
		} else if len(ptr.word) > 0 {
			return ptr.word
		}
		index := word[i] - 'a'
		ptr = ptr.nextNodes[index]
	}
	if ptr == nil {
		return ""
	}
	return ptr.word
}

func (node *trieNode) insert(word string) {
	ptr := node
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if ptr.nextNodes[index] == nil {
			ptr.nextNodes[index] = &trieNode{}
		}
		ptr = ptr.nextNodes[index]
	}
	ptr.word = word
}
