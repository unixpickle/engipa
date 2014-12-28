package engipa

import "strings"

func Guess(word string) Word {
	guesses := "ɑː b k d ɛ f ɡ h ɪ dʒ k l m n ɒ p k r s t oʊ v w ks j z"
	list := strings.Split(guesses, " ")
	res := ""
	for _, rune := range word {
		if rune >= 'A' && rune <= 'Z' {
			rune = 'a' + rune - 'A'
		}
		if rune < 'a' || rune > 'z' {
			continue
		}
		res = res + list[rune-'a']
	}
	return ParseWord(res)
}
