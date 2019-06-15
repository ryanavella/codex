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
	s = stripCommonPrefixes(s)
	if val, ok := titleTab[s]; ok {
		return val, nil
	}
	return UNK, ErrUnknownBook
}

func stripCommonPrefixes(s string) string {
	for _, prefix := range commonPrefixes {
		if strings.HasPrefix(s, prefix) {
			s = s[len(prefix):]
		}
	}
	return s
}

// A list of common book title prefixes. Order matters here.
var commonPrefixes = []string{
	"the letter of paul to the ",
	"letter of paul to the ",
	"letter of paul to ",
	"paul s letter to the ", // '\'' will be replaced with space
	"paul s letter to ",     // '\'' will be replaced with space
	"the gospel according to ",
	"gospel according to ",
	"the letter to the ",
	"the letter of ",
	"the book of the ",
	"the book of ",
	"book of the ",
	"book of ",
	"the ",
}

func nonAlNum(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
