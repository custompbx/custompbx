package mainStruct

import "strings"

const DefaultLocale = "en"

var supportedLocales = map[string]struct{}{
	"en": {}, "fr": {}, "de": {}, "es": {}, "pt-BR": {},
	"it": {}, "tr": {}, "ru": {}, "ar": {}, "fa": {},
	"hi": {}, "zh-Hans": {}, "ja": {}, "ko": {},
}

func NormalizeLocale(locale string) string {
	locale = strings.TrimSpace(locale)
	if _, ok := supportedLocales[locale]; ok {
		return locale
	}
	return DefaultLocale
}

func IsSupportedLocale(locale string) bool {
	_, ok := supportedLocales[strings.TrimSpace(locale)]
	return ok
}

func LegacyLocale(lang uint) string {
	if lang == 1 {
		return "ru"
	}
	return DefaultLocale
}
