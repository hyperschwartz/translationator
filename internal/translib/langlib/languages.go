package langlib

import "golang.org/x/text/language"

// AllLanguages Golang has no magic that allows declared constants to be iterated upon, so an array of this nature must
// be declared in order for the app to continue to pick langlib randomly from the list
var AllLanguages = []language.Tag{
	language.Afrikaans,
	language.Albanian,
	language.AmericanEnglish,
	language.Amharic,
	language.Arabic,
	language.Armenian,
	language.Azerbaijani,
	language.Bengali,
	language.BrazilianPortuguese,
	language.BritishEnglish,
	language.Bulgarian,
	language.Burmese,
	language.CanadianFrench,
	language.Catalan,
	language.Chinese,
	language.Croatian,
	language.Czech,
	language.Danish,
	language.Dutch,
	language.English,
	language.Estonian,
	language.EuropeanPortuguese,
	language.EuropeanSpanish,
	language.Filipino,
	language.French,
	language.Georgian,
	language.German,
	language.Greek,
	language.Gujarati,
	language.Hebrew,
	language.Hindi,
	language.Hungarian,
	language.Icelandic,
	language.Indonesian,
	language.Italian,
	language.Japanese,
	language.Kannada,
	language.Kazakh,
	language.Khmer,
	language.Kirghiz,
	language.Korean,
	language.Lao,
	language.LatinAmericanSpanish,
	language.Latvian,
	language.Lithuanian,
	language.Macedonian,
	language.Malay,
	language.Malayalam,
	language.Marathi,
	language.ModernStandardArabic,
	language.Mongolian,
	language.Nepali,
	language.Norwegian,
	language.Persian,
	language.Polish,
	language.Portuguese,
	language.Punjabi,
	language.Romanian,
	language.Russian,
	language.Serbian,
	language.SerbianLatin,
	language.SimplifiedChinese,
	language.Sinhala,
	language.Slovak,
	language.Slovenian,
	language.Spanish,
	language.Swahili,
	language.Swedish,
	language.Tamil,
	language.Telugu,
	language.Thai,
	language.TraditionalChinese,
	language.Turkish,
	language.Ukrainian,
	language.Urdu,
	language.Uzbek,
	language.Vietnamese,
	language.Zulu,
}

var RandomizerLanguageCodes = FilterLanguageCodes(AllLanguages, func(code language.Tag) bool { return code != language.English })

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
