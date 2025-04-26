package service

import (
	"regexp"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseOrText(s string) (string, error) {
	isText, err := regexp.MatchString(`[А-Яа-я]`, s)
	if err != nil {
		return "", err
	}

	if isText {
		return morse.ToMorse(s), nil
	}

	return morse.ToText(s), nil
}
