package engipa

import "strings"

type Phoneme struct {
	IPA     string
	English string
	Vowel   bool
}

func Alphabet() []Phoneme {
	letters := "ɑː a 1  ɒ o 1  æ a 1  aɪ i 1  aʊ ou 1  ɛ e 1  eɪ a 1  " +
		"ɪ i 1  iː ee 1  ɔː aw 1  ɔɪ oi 1  oʊ o 1  ʊ oo 1  uː oo 1  " +
		"juː ue 1  ʌ u 1  ə a 1  ɨ e 1  ɵ o 1  ʉ u 1  i i 1  b b 0  " +
		"d d 0  ð th 0  dʒ j 0  f f 0  ɡ g 0  h h 0  j y 0  k k 0  " +
		"l l 0  m m 0  n n 0  ŋ ng 0  θ th 0  p p 0  r r 0  s s 0  " +
		"ʃ sh 0  t t 0  tʃ ch 0  v v 0  w w 0  hw wh 0  z z 0  ʒ ti 0  " +
		"x ch 0  ʔ - 0  ʊː oo 1"
	packed := strings.Split(letters, "  ")
	res := make([]Phoneme, len(packed))
	for x, p := range packed {
		comps := strings.Split(p, " ")
		res[x] = Phoneme{comps[0], comps[1], comps[2] == "1"}
	}
	return res
}

func LongestToShortest() []Phoneme {
	alphabet := Alphabet()
	lens := map[int][]Phoneme{1: []Phoneme{}, 2: []Phoneme{}, 3: []Phoneme{}}
	for _, p := range alphabet {
		l := len([]rune(p.IPA))
		lens[l] = append(lens[l], p)
	}
	res := make([]Phoneme, 0, len(alphabet))
	for i := 3; i > 0; i-- {
		res = append(res, lens[i]...)
	}
	return res
}
