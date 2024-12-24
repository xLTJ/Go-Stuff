package romannumerals

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	value       int
	latinSymbol string
}

// Version 2 (1547 ns/op) -------------------------------------------------------------------

var latinToRoman = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts a latin integer into a roman numeral
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("input out of range 1-3999")
	}

	var builder = strings.Builder{}
	for _, latinSymbol := range latinToRoman {
		for input >= latinSymbol.value {
			builder.WriteString(latinSymbol.latinSymbol)
			input -= latinSymbol.value
		}
	}

	return builder.String(), nil
}

// Version 1 (3022 ns/op) ----------------------------------------------------------------

//var latinToRoman = map[int]rune{
//	1000: 'M',
//	500:  'D',
//	100:  'C',
//	50:   'L',
//	10:   'X',
//	5:    'V',
//	1:    'I',
//}
//
//func ToRomanNumeral(input int) (string, error) {
//	if input <= 0 || input >= 4000 {
//		return "", errors.New("input out of range 1-3999")
//	}
//
//	inputString := strconv.Itoa(input)
//	var builder = strings.Builder{}
//	for i, digit := range inputString {
//		newInput, lettersToAdd := addRomanNumbers(input, len(inputString)-i, int(digit-'0'))
//		input = newInput
//		builder.WriteString(lettersToAdd)
//	}
//
//	return builder.String(), nil
//}
//
//func addRomanNumbers(input int, digitsLeft int, currentDigit int) (int, string) {
//	var builder = strings.Builder{}
//	letterTen := latinToRoman[10*int(math.Pow(10, float64(digitsLeft))/10)]
//	letterFive := latinToRoman[5*int(math.Pow(10, float64(digitsLeft))/10)]
//	letterOne := latinToRoman[int(math.Pow(10, float64(digitsLeft))/10)]
//
//	switch {
//	case currentDigit == 9:
//		builder.WriteRune(letterOne)
//		builder.WriteRune(letterTen)
//	case currentDigit >= 5:
//		builder.WriteRune(letterFive)
//		for i := 6; i <= currentDigit; i++ {
//			builder.WriteRune(letterOne)
//		}
//	case currentDigit == 4:
//		builder.WriteRune(letterOne)
//		builder.WriteRune(letterFive)
//	default:
//		for i := 1; i <= currentDigit; i++ {
//			builder.WriteRune(letterOne)
//		}
//	}
//
//	return input % int(math.Pow10(digitsLeft-1)), builder.String()
//}
