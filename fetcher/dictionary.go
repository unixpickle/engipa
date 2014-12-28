package fetcher

import "strings"

func DictionaryIPA(word string) (string, string, error) {
	url := "http://dictionary.reference.com/browse/" + word
	body, err := FetchPage(url)
	if err != nil {
		return "", "", err
	}
	ipa, err := FindTag(body, "<span class=\"pron ipapron\">", " </span>")
	if err != nil {
		return "", "", err
	}
	heading, err := FindTag(body, "<span class=\"me\" data-syllable=\".*?\">",
		"</span>")
	if err != nil {
		return "", "", err
	}
	idx := strings.Index(ipa, ",")
	if idx >= 0 {
		ipa = ipa[0:idx]
	}
	ipa = strings.Replace(ipa, "y", "j", -1)
	ipa = strings.Replace(strings.Replace(ipa, "ɑː", "ɑ", -1), "ɑ", "ɑː",
		-1)
	return ipa, heading, nil
}
