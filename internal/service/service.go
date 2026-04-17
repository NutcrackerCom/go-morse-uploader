package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("empty input")
	}

	for _, r := range s {
		switch r {
		case '.', '-', ' ', '/':
		default:
			return morse.ToMorse(s), nil
		}
	}
	return morse.ToText(s), nil
}
