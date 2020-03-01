package generator

import "github.com/elahe-dastan/urlShortener_KGS/model"

const LengthOfShortURL = 2

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateURLsRec(prefix string, k int, shortURLs *[]model.ShortURL) {
	if k == 0 {
		*shortURLs = append(*shortURLs, model.ShortURL{URL: prefix, IsUsed: false})
		return
	}

	k--

	for i := range source {
		newPrefix := prefix + string(source[i])
		generateURLsRec(newPrefix, k, shortURLs)
	}
}

func Generate() []model.ShortURL {
	var shortURLs []model.ShortURL

	generateURLsRec("", LengthOfShortURL, &shortURLs)

	return shortURLs
}
