package generator

import "github.com/elahe-dastan/urlShortener_KGS/model"

const LengthOfShortURL  =  2

const source  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var shortURLs []model.ShortURL

func generateURLsRec(prefix string,k int )  {
	if k == 0 {
		shortURLs = append(shortURLs, model.ShortURL{URL: prefix, IsUsed:false})
		return
	}

	k--
	for i := range source {
		newPrefix := prefix + string(source[i])
		generateURLsRec(newPrefix, k)
	}
}

func Generate()  []model.ShortURL {
	generateURLsRec("", LengthOfShortURL)
	return shortURLs
}

