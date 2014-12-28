package fetcher

import (
	"errors"
	"github.com/unixpickle/engipa"
	"strings"
)

func Cleanup(word string) string {
	return engipa.ParseWord(word).String()
}

func GuessIPA(word string) (string, error) {
	guess, exact, err := DictionaryIPA(word)
	if err == nil && exact == word {
		return Cleanup(guess), nil
	}
	guess1, exact1, err1 := CambridgeIPA(word)
	if err1 == nil && exact1 == word {
		return Cleanup(guess1), nil
	}
	if err1 != nil && err != nil {
		return "", errors.New("Failed to lookup word.")
	}
	if err == nil {
		return GenerateGuess(word, exact, guess)
	} else {
		return GenerateGuess(word, exact1, guess1)
	}
}

func GenerateGuess(word, exact, ipa string) (string, error) {
	// This is my modest attempt at adding common suffixes to words.
	parsed := engipa.ParseWord(ipa)
	ipa = parsed.String()
	if strings.HasSuffix(word, "ings") {
		ipa = ipa + "ɪŋs"
	} else if strings.HasSuffix(word, "ing") {
		ipa = ipa + "ɪŋ"
	} else if strings.HasSuffix(word, "ly") {
		ipa = ipa + "lē"
	} else if strings.HasSuffix(word, "d") && len(parsed) > 0 {
		last := parsed[len(parsed)-1].IPA
		if (!parsed.EndsVowel() || last == "n") && last != "t" && last != "r" {
			ipa = ipa + "t"
		} else {
			ipa = ipa + "d"
		}
	} else if strings.HasSuffix(word, "rs") {
		ipa = ipa + "ərs"
	} else if strings.HasSuffix(word, "r") {
		ipa = ipa + "ər"
	} else if strings.HasSuffix(word, "est") {
		ipa = ipa + "əst"
	} else if strings.HasSuffix(word, "ments") {
		ipa = ipa + "mənts"
	} else if strings.HasSuffix(word, "ment") {
		ipa = ipa + "mənt"
	} else if strings.HasSuffix(word, "cy") {
		ipa = ipa + "si"
	} else if strings.HasSuffix(word, "cies") {
		ipa = ipa + "siz"
	} else if strings.HasSuffix(word, "s") && len(parsed) > 0 {
		last := parsed[len(parsed)-1].IPA
		if last == "s" || last == "z" || last == "ʃ" || last == "tʃ" {
			ipa = ipa + "ez"
		} else {
			ipa = ipa + "z"
		}
	} else if strings.HasSuffix(word, "able") {
		ipa = ipa + "əbəl"
	} else if strings.HasSuffix(word, "ship") {
		ipa = ipa + "ʃɪp"
	} else if strings.HasSuffix(word, "ary") {
		ipa = ipa + "əri"
	} else if strings.HasSuffix(word, "atory") {
		ipa = ipa + "ətɔri"
	} else if strings.HasSuffix(word, "ory") {
		ipa = ipa + "ɔri"
	} else if strings.HasSuffix(word, "al") {
		ipa = ipa + "ʊl"
	} else if strings.HasSuffix(word, "ive") {
		ipa = ipa + "ɪv"
	} else {
		return "", errors.New("Unknown suffix.")
	}
	return ipa, nil
}

