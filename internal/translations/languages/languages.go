package languages

import "golang.org/x/text/language"

var AllLanguages = []language.Tag{
	language.Afrikaans,
	language.Amharic,
	language.Arabic,
	language.ModernStandardArabic,
	language.Azerbaijani,
	language.Bulgarian,
	language.Bengali,
	language.Catalan,
	language.Czech,
	language.Danish,
	language.German,
	language.Greek,
	language.English,
	language.AmericanEnglish,
	language.BritishEnglish,
	language.Spanish,
	language.EuropeanSpanish,
	language.LatinAmericanSpanish,
	language.Estonian,
	language.Persian,
	language.Finnish,
	language.Filipino,
	language.French,
	language.CanadianFrench,
	language.Gujarati,
	language.Hebrew,
	language.Hindi,
	language.Croatian,
	language.Hungarian,
	language.Armenian,
	language.Indonesian,
	language.Icelandic,
	language.Italian,
	language.Japanese,
	language.Georgian,
	language.Kazakh,
	language.Khmer,
	language.Kannada,
	language.Korean,
	language.Kirghiz,
	language.Lao,
	language.Lithuanian,
	language.Latvian,
	language.Macedonian,
	language.Malayalam,
	language.Mongolian,
	language.Marathi,
	language.Malay,
	language.Burmese,
	language.Nepali,
	language.Dutch,
	language.Norwegian,
	language.Punjabi,
	language.Polish,
	language.Portuguese,
	language.BrazilianPortuguese,
	language.EuropeanPortuguese,
	language.Romanian,
	language.Russian,
	language.Sinhala,
	language.Slovak,
	language.Slovenian,
	language.Albanian,
	language.Serbian,
	language.SerbianLatin,
	language.Swedish,
	language.Swahili,
	language.Tamil,
	language.Telugu,
	language.Thai,
	language.Turkish,
	language.Ukrainian,
	language.Urdu,
	language.Uzbek,
	language.Vietnamese,
	language.Chinese,
	language.SimplifiedChinese,
	language.TraditionalChinese,
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
