package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseOrText(s string) (string, error) {
	f := func(r rune) bool {
		return r == '.' || r == '-'
	}

	isMorse := strings.ContainsFunc(s, f)
	if isMorse {
		return morse.ToText(s), nil
	}

	return morse.ToMorse(s), nil
}
