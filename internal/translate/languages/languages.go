package languages

type LanguageCode string

const (
	Undefined          LanguageCode = "UNDEFINED"
	Afrikaans                       = "af"
	Albanian                        = "sq"
	Amharic                         = "am"
	Arabic                          = "ar"
	Armenian                        = "hy"
	Azeerbaijani                    = "az"
	Bashkir                         = "ba"
	Basque                          = "eu"
	Belarusian                      = "be"
	Bengali                         = "bn"
	Bosnian                         = "bs"
	Bulgarian                       = "bg"
	Burmese                         = "my"
	Catalan                         = "ca"
	Cebuano                         = "ceb"
	Chichewa                        = "ny"
	ChineseSimplified               = "zh-CN"
	ChineseTraditional              = "zh-TW"
	Corsican                        = "co"
	Croatian                        = "hr"
	Czech                           = "cs"
	Danish                          = "da"
	Dutch                           = "nl"
	English                         = "en"
	Esperanto                       = "eo"
	Estonian                        = "et"
	Finnish                         = "fi"
	French                          = "fr"
	Frisian                         = "fy"
	Galician                        = "gl"
	Georgian                        = "ka"
	German                          = "de"
	Greek                           = "el"
	Gujarati                        = "gu"
	HaitianCreole                   = "ht"
	Hausa                           = "ha"
	Hawaiian                        = "haw"
	Hebrew                          = "iw"
	Hill_Mari                       = "mrj"
	Hindi                           = "hi"
	Hmong                           = "hmn"
	Hungarian                       = "hu"
	Icelandic                       = "is"
	Igbo                            = "ig"
	Indonesian                      = "id"
	Irish                           = "ga"
	Italian                         = "it"
	Japanese                        = "ja"
	Javanese                        = "jw"
	Kannada                         = "kn"
	Kazakh                          = "kk"
	Khmer                           = "km"
	Korean                          = "ko"
	Kurdish                         = "ku"
	Kyrgyz                          = "ky"
	Lao                             = "lo"
	Latin                           = "la"
	Latvian                         = "lv"
	Lithuanian                      = "lt"
	Luxembourgish                   = "lb"
	Macedonian                      = "mk"
	Malagasy                        = "mg"
	Malay                           = "ms"
	Malayalam                       = "ml"
	Maltese                         = "mt"
	Maori                           = "mi"
	Marathi                         = "mr"
	Mari                            = "mhr"
	Mongolian                       = "mn"
	Nepali                          = "ne"
	Norwegian                       = "no"
	Pashto                          = "ps"
	Papiamento                      = "pap"
	Persian                         = "fa"
	Polish                          = "pl"
	Portuguese                      = "pt"
	Punjabi                         = "pa"
	Romanian                        = "ro"
	Russian                         = "ru"
	Samoan                          = "sm"
	ScotsGaelic                     = "gd"
	Serbian                         = "sr"
	Sesotho                         = "st"
	Shona                           = "sn"
	Sindhi                          = "sd"
	Sinhala                         = "si"
	Slovak                          = "sk"
	Slovenian                       = "sl"
	Somali                          = "so"
	Spanish                         = "es"
	Sundanese                       = "su"
	Swahili                         = "sw"
	Swedish                         = "sv"
	TagalogFilipino                 = "tl"
	Tajik                           = "tg"
	Tamil                           = "ta"
	Tatar                           = "tt"
	Telugu                          = "te"
	Thai                            = "th"
	Turkish                         = "tr"
	Udmurt                          = "udm"
	Ukrainian                       = "uk"
	Urdu                            = "ur"
	Uzbek                           = "uz"
	Vietnamese                      = "vi"
	Welsh                           = "cy"
	Xhosa                           = "xh"
	Yiddish                         = "yi"
	Yoruba                          = "yo"
	Zulu                            = "zu"
)

var LanguageCodes = []LanguageCode{
	Undefined,
	Afrikaans,
	Albanian,
	Amharic,
	Arabic,
	Armenian,
	Azeerbaijani,
	Bashkir,
	Basque,
	Belarusian,
	Bengali,
	Bosnian,
	Bulgarian,
	Burmese,
	Catalan,
	Cebuano,
	Chichewa,
	ChineseSimplified,
	ChineseTraditional,
	Corsican,
	Croatian,
	Czech,
	Danish,
	Dutch,
	English,
	Esperanto,
	Estonian,
	Finnish,
	French,
	Frisian,
	Galician,
	Georgian,
	German,
	Greek,
	Gujarati,
	HaitianCreole,
	Hausa,
	Hawaiian,
	Hebrew,
	Hill_Mari,
	Hindi,
	Hmong,
	Hungarian,
	Icelandic,
	Igbo,
	Indonesian,
	Irish,
	Italian,
	Japanese,
	Javanese,
	Kannada,
	Kazakh,
	Khmer,
	Korean,
	Kurdish,
	Kyrgyz,
	Lao,
	Latin,
	Latvian,
	Lithuanian,
	Luxembourgish,
	Macedonian,
	Malagasy,
	Malay,
	Malayalam,
	Maltese,
	Maori,
	Marathi,
	Mari,
	Mongolian,
	Nepali,
	Norwegian,
	Pashto,
	Papiamento,
	Persian,
	Polish,
	Portuguese,
	Punjabi,
	Romanian,
	Russian,
	Samoan,
	ScotsGaelic,
	Serbian,
	Sesotho,
	Shona,
	Sindhi,
	Sinhala,
	Slovak,
	Slovenian,
	Somali,
	Spanish,
	Sundanese,
	Swahili,
	Swedish,
	TagalogFilipino,
	Tajik,
	Tamil,
	Tatar,
	Telugu,
	Thai,
	Turkish,
	Udmurt,
	Ukrainian,
	Urdu,
	Uzbek,
	Vietnamese,
	Welsh,
	Xhosa,
	Yiddish,
	Yoruba,
	Zulu,
}

var RandomizerLanguageCodes = filterLanguageCodes(LanguageCodes, func(code LanguageCode) bool { return code != English && code != Undefined })

func filterLanguageCodes(array []LanguageCode, filterFn func(LanguageCode) bool) (result []LanguageCode) {
	for _, val := range array {
		if filterFn(val) {
			result = append(result, val)
		}
	}
	return
}
