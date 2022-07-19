package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var (
		sb               strings.Builder
		multiplier       int
		addStr           string
		backSlashCounter = 0
	)

	sRune := []rune(s)

	for i, curRune := range sRune {
		switch {
		case isBackSlash(curRune):
			backSlashCounter++
			if backSlashCounter%2 == 0 {
				addStr = string(curRune)
			} else {
				continue
			}
		case isFirstDigit(curRune, i):
			return "", ErrInvalidString
		case isDoubleDigit(curRune, i, s):
			return "", ErrInvalidString
		case unicode.IsDigit(curRune):
			digit, err := strconv.Atoi(string(curRune))
			if err != nil {
				return "", err
			}

			if backSlashCounter%2 == 1 {
				addStr = string(curRune)
				multiplier = 1
			} else {
				// case \n5
				if isEscapedChar(i, s) {
					addStr = "\\" + string(getPrevRune(i, s))
					multiplier = digit
				} else {
					// case n5
					addStr = string(getPrevRune(i, s))
					multiplier = digit - 1
				}
			}
			backSlashCounter = 0
		case isBackSlash(getPrevRune(i, s)):
			backSlashCounter = 0
			continue
		default:
			multiplier = 1
			backSlashCounter = 0
			addStr = string(curRune)
		}

		addString(&sb, addStr, multiplier)

		fmt.Println(sb.String())
	}

	return sb.String(), nil
}

func isEscapedChar(i int, s string) bool {
	return isBackSlash(getRune(i-2, s)) && !unicode.IsDigit(getPrevRune(i, s)) &&
		!isBackSlash(getPrevRune(i, s))
}

func isDoubleDigit(curRune int32, i int, s string) bool {
	return unicode.IsDigit(curRune) && !isBackSlash(getPrevRune(i, s)) && isNextDigit(i+1, s)
}

func isFirstDigit(curRune int32, i int) bool {
	return unicode.IsDigit(curRune) && i == 0
}

func getRune(i int, s string) rune {
	if i < 0 {
		return 0
	}

	sT := []rune(s)

	if i >= len(sT) {
		return 0
	}

	return sT[i]
}

func getPrevRune(i int, s string) rune {
	return getRune(i-1, s)
}

func isNextDigit(i int, s string) bool {
	if i < 0 {
		return false
	}

	sT := []rune(s)

	if i >= len(sT) {
		return false
	}

	return unicode.IsDigit(sT[i])
}

func isBackSlash(r rune) bool {
	return r == 92
}

func addString(sb *strings.Builder, s string, m int) {
	if m <= 0 {
		t := sb.String()
		sb.Reset()
		sT := []rune(t)
		sT = sT[:len(sT)-1]
		sb.WriteString(string(sT))
	} else {
		sb.WriteString(strings.Repeat(s, m))
	}
}
