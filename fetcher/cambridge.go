package fetcher

import "strings"

func CambridgeIPA(word string) (string, string, error) {
	url := "http://dictionary.cambridge.org/us/search/american-english/" +
		"direct/?q=" + word
	body, err := FetchPage(url)
	if err != nil {
		return "", "", err
	}
	ipa, err := FindTag(body, "<span class=\"ipa\">", "</span>")
	if err != nil {
		return "", "", err
	}
	heading, err := FindTag(body,
		"<h2 class=\"di-title cdo-section-title-hw\">", "</h2>")
	if err != nil {
		return "", "", err
	}
	ipa = strings.Replace(ipa, "e", "ɛ", -1)
	return ipa, heading, nil
}
