package luhn

// Valid checks if an id is valid per the Luhn formula (313.4 ns/op)
func Valid(id string) bool {
	counter := 0 // using a counter instead of checking length with len() is actually slightly faster
	// (and it also doesnt get fucked by single digits with spaces which was why i originally tried it lol)
	sum := 0
	shouldDouble := false
	for i := len(id) - 1; i >= 0; i-- {
		character := id[i]
		if character == ' ' {
			continue
		}

		var digit int
		if character >= '0' && character <= '9' {
			digit = int(character - '0')
		} else {
			return false
		}

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
		counter++
	}

	return sum%10 == 0 && counter > 1
}

// 53795 ns/op
//func Valid(id string) bool {
//	id = strings.ReplaceAll(id, " ", "")
//	onlyNumbersFormat := regexp.MustCompile(`^[0-9]+$`)
//
//	if len(id) <= 1 || !onlyNumbersFormat.MatchString(id) {
//		return false
//	}
//
//	startIndex := len(id) % 2
//	sum := 0
//	for i, numberRune := range id {
//		number := int(numberRune - '0')
//
//		if i%2 == startIndex {
//			number *= 2
//		}
//		if number > 9 {
//			number -= 9
//		}
//		sum += number
//	}
//
//	return sum%10 == 0
//}
