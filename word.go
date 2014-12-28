package engipa

import "strings"

type Word []Phoneme

// ParseWord turns an IPA string into a phonetic word.
func ParseWord(word string) Word {
	if len(word) == 0 {
		return Word{}
	}
	alphabet := LongestToShortest()
	for _, p := range alphabet {
		if strings.HasPrefix(word, p.IPA) {
			res := []Phoneme{p}
			rem := word[len(p.IPA):]
			return append(res, ParseWord(rem)...)
		}
	}
	// Skip the first rune
	tail := string([]rune(word)[1:])
	return ParseWord(tail)
}

func (w Word) StartsVowel() bool {
	if len(w) == 0 {
		return false
	}
	return w[0].Vowel
}

func (w Word) EndsVowel() bool {
	if len(w) == 0 {
		return false
	}
	return w[len(w)-1].Vowel
}

func (w Word) String() string {
	res := ""
	for _, p := range w {
		res += p.IPA
	}
	return res
}
