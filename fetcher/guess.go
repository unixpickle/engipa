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
	guess, exact, err := CambridgeIPA(word)
	if err == nil && exact == word {
		return Cleanup(guess), nil
	}
	guess1, exact1, err1 := DictionaryIPA(word)
	if err1 == nil && exact1 == word {
		return Cleanup(guess1), nil
	}
	if err1 != nil && err != nil {
		return "", errors.New("Failed to lookup word.")
	}
	if err == nil {
		return GenerateGuess(word, exact, guess), nil
	} else {
		return GenerateGuess(word, exact1, guess1), nil
	}
}

func GenerateGuess(word, exact, ipa string) string {
	// This is my modest attempt at adding common suffixes to words.
	parsed := engipa.ParseWord(ipa)
	ipa = parsed.String()
	if strings.HasSuffix(word, "ing") {
		ipa = ipa + "ɪŋ"
	} else if strings.HasSuffix(word, "ly") {
		ipa = ipa + "lē"
	} else if strings.HasSuffix(word, "d") {
		if !parsed.EndsVowel() {
			ipa = ipa + "t"
		} else {
			ipa = ipa + "d"
		}
	} else if strings.HasSuffix(word, "r") {
		ipa = ipa + "ər"
	} else if strings.HasSuffix(word, "est") {
		ipa = ipa + "əst"
	}
	return ipa
}
