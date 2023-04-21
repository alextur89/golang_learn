package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (result string, err error) {
	var prev_char, current_char rune
	var index int
	for index, current_char = range s {
		if unicode.IsNumber(current_char) {
			if index == 0 || unicode.IsNumber(prev_char) {
				err = ErrInvalidString
				result = ""
				return
			}
			result += strings.Repeat(string(prev_char), int(current_char-'0'))
		} else if index > 0 && !unicode.IsNumber(prev_char) {
			result += string(prev_char)
		}
		prev_char = current_char
	}
	if index > 0 && !unicode.IsNumber(prev_char) {
		result += string(prev_char)
	}
	return
}
