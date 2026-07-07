package logsafe

import (
	"encoding/json"
	"fmt"
	"regexp"
)

const redacted = "[REDACTED]"

var (
	jsonSecretField = regexp.MustCompile(`(?i)("([^"]*(token|pass|password|secret|authorization|authval|nodepw)[^"]*)"\s*:\s*)("[^"]*"|null|[0-9]+|true|false)`)
	dsnPassword     = regexp.MustCompile(`(?i)(password=)[^\s]+`)
	urlPassword     = regexp.MustCompile(`(?i)(://[^:\s/@]+:)[^@\s]+(@)`)
)

// Redact returns a log-safe representation of value with credential-bearing
// fields masked. It is intended for diagnostics only, not for persistence.
func Redact(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return RedactString(v)
	case fmt.Stringer:
		return RedactString(v.String())
	default:
		raw, err := json.Marshal(value)
		if err != nil {
			return RedactString(fmt.Sprintf("%+v", value))
		}
		return RedactString(string(raw))
	}
}

func RedactString(value string) string {
	value = jsonSecretField.ReplaceAllString(value, `${1}"`+redacted+`"`)
	value = dsnPassword.ReplaceAllString(value, `${1}`+redacted)
	value = urlPassword.ReplaceAllString(value, `${1}`+redacted+`${2}`)
	return value
}
