package hfuncs

import (
	"strings"
)

func Trimmer(str string) string {
	return strings.Trim(str, "\"`'!@%#_-{}|[]:;.")
}
