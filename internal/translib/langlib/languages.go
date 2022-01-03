package langlib

import (
	log "github.com/siruspen/logrus"
	"golang.org/x/text/language"
)

// validLanguageCodes Derived via: https://cloud.google.com/translate/docs/languages
// Each code contained in this list will be accepted via the Google Translate api, and will allow many subsequent
// requests to be made without fail.
var validLanguageCodes = []string{
	"af",
	"sq",
	"am",
	"ar",
	"hy",
	"az",
	"eu",
	"be",
	"bn",
	"bs",
	"bg",
	"ca",
	"ceb",
	"zh",
	"zh-TW",
	"co",
	"hr",
	"cs",
	"da",
	"nl",
	"en",
	"eo",
	"et",
	"fi",
	"fr",
	"fy",
	"gl",
	"ka",
	"de",
	"el",
	"gu",
	"ht",
	"ha",
	"haw",
	"he",
	"hi",
	"hmn",
	"hu",
	"is",
	"ig",
	"id",
	"ga",
	"it",
	"ja",
	"jv",
	"kn",
	"kk",
	"km",
	"rw",
	"ko",
	"ku",
	"ky",
	"lo",
	"lv",
	"lt",
	"lb",
	"mk",
	"mg",
	"ms",
	"ml",
	"mt",
	"mi",
	"mr",
	"mn",
	"my",
	"ne",
	"no",
	"ny",
	"or",
	"ps",
	"fa",
	"pl",
	"pt",
	"pa",
	"ro",
	"ru",
	"sm",
	"gd",
	"sr",
	"st",
	"sn",
	"sd",
	"si",
	"sk",
	"sl",
	"so",
	"es",
	"su",
	"sw",
	"sv",
	"tl",
	"tg",
	"ta",
	"tt",
	"te",
	"th",
	"tr",
	"tk",
	"uk",
	"ur",
	"ug",
	"uz",
	"vi",
	"cy",
	"xh",
	"yi",
	"yo",
	"zu",
}

var AllValidLanguages = parseAllValidLanguages()

func parseAllValidLanguages() (validLanguages []language.Tag) {
	for _, code := range validLanguageCodes {
		tag, err := language.Parse(code)
		if err != nil {
			log.Error("Invalid language code provided: ", code, ". Please remove it. Error: ", err)
		}
		validLanguages = append(validLanguages, tag)
	}
	return validLanguages
}

var RandomizerLanguageCodes = FilterLanguageCodes(AllValidLanguages, func(code language.Tag) bool { return code != language.English })

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
