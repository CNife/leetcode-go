package implement_magic_dictionary

type MagicDictionary struct {
	dictionary []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (md *MagicDictionary) BuildDict(dictionary []string) {
	md.dictionary = dictionary
}

func (md *MagicDictionary) Search(word string) bool {
	for _, dictWord := range md.dictionary {
		if len(word) == len(dictWord) {
			replaced := 0
			for i := 0; i < len(word); i++ {
				if word[i] != dictWord[i] {
					replaced++
				}
			}
			if replaced == 1 {
				return true
			}
		}
	}
	return false
}
