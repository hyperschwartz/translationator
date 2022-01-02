package languages

type LanguageCode string

const (
	Undefined          LanguageCode = "UNDEFINED"
	Afrikaans          LanguageCode = "af"
	Albanian           LanguageCode = "sq"
	Amharic            LanguageCode = "am"
	Arabic             LanguageCode = "ar"
	Armenian           LanguageCode = "hy"
	Azeerbaijani       LanguageCode = "az"
	Bashkir            LanguageCode = "ba"
	Basque             LanguageCode = "eu"
	Belarusian         LanguageCode = "be"
	Bengali            LanguageCode = "bn"
	Bosnian            LanguageCode = "bs"
	Bulgarian          LanguageCode = "bg"
	Burmese            LanguageCode = "my"
	Catalan            LanguageCode = "ca"
	Cebuano            LanguageCode = "ceb"
	Chichewa           LanguageCode = "ny"
	ChineseSimplified  LanguageCode = "zh-CN"
	ChineseTraditional LanguageCode = "zh-TW"
	Corsican           LanguageCode = "co"
	Croatian           LanguageCode = "hr"
	Czech              LanguageCode = "cs"
	Danish             LanguageCode = "da"
	Dutch              LanguageCode = "nl"
	English            LanguageCode = "en"
	Esperanto          LanguageCode = "eo"
	Estonian           LanguageCode = "et"
	Finnish            LanguageCode = "fi"
	French             LanguageCode = "fr"
	Frisian            LanguageCode = "fy"
	Galician           LanguageCode = "gl"
	Georgian           LanguageCode = "ka"
	German             LanguageCode = "de"
	Greek              LanguageCode = "el"
	Gujarati           LanguageCode = "gu"
	HaitianCreole      LanguageCode = "ht"
	Hausa              LanguageCode = "ha"
	Hawaiian           LanguageCode = "haw"
	Hebrew             LanguageCode = "iw"
	Hill_Mari          LanguageCode = "mrj"
	Hindi              LanguageCode = "hi"
	Hmong              LanguageCode = "hmn"
	Hungarian          LanguageCode = "hu"
	Icelandic          LanguageCode = "is"
	Igbo               LanguageCode = "ig"
	Indonesian         LanguageCode = "id"
	Irish              LanguageCode = "ga"
	Italian            LanguageCode = "it"
	Japanese           LanguageCode = "ja"
	Javanese           LanguageCode = "jw"
	Kannada            LanguageCode = "kn"
	Kazakh             LanguageCode = "kk"
	Khmer              LanguageCode = "km"
	Korean             LanguageCode = "ko"
	Kurdish            LanguageCode = "ku"
	Kyrgyz             LanguageCode = "ky"
	Lao                LanguageCode = "lo"
	Latin              LanguageCode = "la"
	Latvian            LanguageCode = "lv"
	Lithuanian         LanguageCode = "lt"
	Luxembourgish      LanguageCode = "lb"
	Macedonian         LanguageCode = "mk"
	Malagasy           LanguageCode = "mg"
	Malay              LanguageCode = "ms"
	Malayalam          LanguageCode = "ml"
	Maltese            LanguageCode = "mt"
	Maori              LanguageCode = "mi"
	Marathi            LanguageCode = "mr"
	Mari               LanguageCode = "mhr"
	Mongolian          LanguageCode = "mn"
	Nepali             LanguageCode = "ne"
	Norwegian          LanguageCode = "no"
	Pashto             LanguageCode = "ps"
	Papiamento         LanguageCode = "pap"
	Persian            LanguageCode = "fa"
	Polish             LanguageCode = "pl"
	Portuguese         LanguageCode = "pt"
	Punjabi            LanguageCode = "pa"
	Romanian           LanguageCode = "ro"
	Russian            LanguageCode = "ru"
	Samoan             LanguageCode = "sm"
	ScotsGaelic        LanguageCode = "gd"
	Serbian            LanguageCode = "sr"
	Sesotho            LanguageCode = "st"
	Shona              LanguageCode = "sn"
	Sindhi             LanguageCode = "sd"
	Sinhala            LanguageCode = "si"
	Slovak             LanguageCode = "sk"
	Slovenian          LanguageCode = "sl"
	Somali             LanguageCode = "so"
	Spanish            LanguageCode = "es"
	Sundanese          LanguageCode = "su"
	Swahili            LanguageCode = "sw"
	Swedish            LanguageCode = "sv"
	TagalogFilipino    LanguageCode = "tl"
	Tajik              LanguageCode = "tg"
	Tamil              LanguageCode = "ta"
	Tatar              LanguageCode = "tt"
	Telugu             LanguageCode = "te"
	Thai               LanguageCode = "th"
	Turkish            LanguageCode = "tr"
	Udmurt             LanguageCode = "udm"
	Ukrainian          LanguageCode = "uk"
	Urdu               LanguageCode = "ur"
	Uzbek              LanguageCode = "uz"
	Vietnamese         LanguageCode = "vi"
	Welsh              LanguageCode = "cy"
	Xhosa              LanguageCode = "xh"
	Yiddish            LanguageCode = "yi"
	Yoruba             LanguageCode = "yo"
	Zulu               LanguageCode = "zu"
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

var RandomizerLanguageCodes = FilterLanguageCodes(LanguageCodes, func(code LanguageCode) bool { return code != English && code != Undefined })

func FilterLanguageCodes(array []LanguageCode, filterFn func(LanguageCode) bool) (result []LanguageCode) {
	for _, val := range array {
		if filterFn(val) {
			result = append(result, val)
		}
	}
	return
}
