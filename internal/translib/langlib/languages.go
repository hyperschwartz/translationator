package langlib

import (
	"fmt"
	"golang.org/x/text/language"
)

var AllValidLanguages = parseAllValidLanguages()
var RandomizerLanguageCodes = FilterLanguageCodes(AllValidLanguages, func(code language.Tag) bool { return code != language.English })

// validLanguageCodes Derived via: https://cloud.google.com/translate/docs/languages
// Each code contained in this list will be accepted via the Google Translate api, and will allow many subsequent
// requests to be made without fail.
var validLanguageCodes = []string{
	"af",
	"am",
	"ar",
	"az",
	"be",
	"bg",
	"bn",
	"bs",
	"ca",
	"ceb",
	"co",
	"cs",
	"cy",
	"da",
	"de",
	"el",
	"en",
	"eo",
	"es",
	"et",
	"eu",
	"fa",
	"fi",
	"fr",
	"fy",
	"ga",
	"gd",
	"gl",
	"gu",
	"ha",
	"haw",
	"he",
	"hi",
	"hmn",
	"hr",
	"ht",
	"hu",
	"hy",
	"id",
	"ig",
	"is",
	"it",
	"ja",
	"jv",
	"ka",
	"kk",
	"km",
	"kn",
	"ko",
	"ku",
	"ky",
	"lb",
	"lo",
	"lt",
	"lv",
	"mg",
	"mi",
	"mk",
	"ml",
	"mn",
	"mr",
	"ms",
	"mt",
	"my",
	"ne",
	"nl",
	"no",
	"ny",
	"or",
	"pa",
	"pl",
	"ps",
	"pt",
	"ro",
	"ru",
	"rw",
	"sd",
	"si",
	"sk",
	"sl",
	"sm",
	"sn",
	"so",
	"sq",
	"sr",
	"st",
	"su",
	"sv",
	"sw",
	"ta",
	"te",
	"tg",
	"th",
	"tk",
	"tl",
	"tr",
	"tt",
	"ug",
	"uk",
	"ur",
	"uz",
	"vi",
	"xh",
	"yi",
	"yo",
	"zh",
	"zh-TW",
	"zu",
}

func FilterLanguageCodes(array []language.Tag, filterFn func(language.Tag) bool) (result []language.Tag) {
	for _, val := range array {
		if filterFn(val) {
			result = append(result, val)
		}
	}
	return result
}

func LanguageCodeInArray(array []language.Tag, code language.Tag) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == code {
			return true
		}
	}
	return false
}

func parseAllValidLanguages() (validLanguages []language.Tag) {
	for _, code := range validLanguageCodes {
		tag, err := language.Parse(code)
		if err != nil {
			fmt.Printf("Invalid language code provided: %s. Please remove it. Error: %v", code, err)
		}
		validLanguages = append(validLanguages, tag)
	}
	return validLanguages
}
