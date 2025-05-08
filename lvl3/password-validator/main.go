package main

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(ValidatePassword("qwe"))
	fmt.Println(ValidatePassword("qwertyui"))
	fmt.Println(ValidatePassword("Qwertyui"))
	fmt.Println(ValidatePassword("Qwertyui1"))
}

var LengthErr = errors.New("password must be at least 8 characters long")
var CapitalErr = errors.New("password must contain at least 1 capital char")
var DigitErr = errors.New("password must contain at least 1 digit")

func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return LengthErr
	}

	isCapital := false
	isDigit := false

	for _, char := range pass {
		asciiChar := int(char)

		if !isCapital && unicode.IsUpper(char) {
			isCapital = true
		}

		if !isDigit && asciiChar >= 48 && asciiChar <= 57 {
			isDigit = true
		}
	}

	if !isCapital {
		return CapitalErr
	}

	if !isDigit {
		return DigitErr
	}

	return nil
}
