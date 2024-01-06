package hashtable

import (
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {

	romans := strings.Split(s, "")
	numb := 0

	for i := 0; i < len(romans)-1; i++ {
		current, next := roman[romans[i]], roman[romans[i+1]]

		if current < next {
			numb -= current
		} else {
			numb += current
		}
	}

	numb += roman[romans[len(romans)-1]]

	return numb
}
