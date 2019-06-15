package book

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func parse(s string) (ID, error) {
	s = strings.Join(strings.FieldsFunc(s, nonAlNum), " ")
	s = strings.ToLower(s)
	s = norm.NFC.String(s)
	if val, ok := titleTab[s]; ok {
		return val, nil
	}
	return UNK, ErrUnknownBook
}

func nonAlNum(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
