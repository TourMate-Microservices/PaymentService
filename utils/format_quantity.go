package utils

import (
	"fmt"
	"strconv"
)

func NumberToStringFormat(number int) string {
	var strNumber string = fmt.Sprint(number)
	if number < 1000 {
		return strNumber
	}

	var numberToDivide int = 1
	for i := 1; i <= len(strNumber); i += 3 {
		numberToDivide *= 1000
	}

	var preNumber int = number / numberToDivide
	var sepChar string = "."

	// Generate re number
	var reNumberLength int = len(strNumber) - len(fmt.Sprint(preNumber))
	var reNumber int = number % numberToDivide
	var reNumberStr string = fmt.Sprint(reNumber)

	var gapStrLength int = reNumberLength - len(reNumberStr)
	tmp, _ := strconv.ParseInt(string(reNumberStr[0]), 10, 0)
	var reRes string

	switch gapStrLength {
	case 1:
		tmpNext, _ := strconv.ParseInt(string(reNumberStr[1]), 10, 0)
		if tmpNext >= 8 {
			tmp += 1
		}

		if tmp == 10 {
			preNumber += 1
		} else {
			reRes += "," + fmt.Sprint(tmp) + "k"
		}
	case 2:
		if tmp >= 9 {
			reRes += ",1K"
		}
	}

	// Generate pre number
	var preNumberStr string = fmt.Sprint(preNumber)
	var beginCharsAmount int = len(preNumberStr) % 3
	var preRes string

	switch beginCharsAmount {
	case 1:
		preRes = string(preNumberStr[0])
	case 2:
		preRes = string(preNumberStr[0]) + string(preNumberStr[1])
	}

	if len(preNumberStr) > 3 {
		for i := beginCharsAmount; i < len(preNumberStr); i += 3 {
			preRes += sepChar + string(preNumberStr[i]) + string(preNumberStr[i+1]) + string(preNumberStr[i+2])
		}
	}
	return preRes + reRes
}
