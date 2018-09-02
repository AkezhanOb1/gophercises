package main

import (
	"fmt"
	)

func main() {

	fmt.Println(Encrypter(11,"www.abc.xy", 87))
}

func Encrypter(n int, text string, rotate int) string {

	var cipherText string
	for _, letter := range text {

		if letter > 96 && letter < 123 {
			cipherText += string(encryptSmallLetter(letter, rotate))
			continue
		}else if letter> 64 && letter < 91 {
			cipherText += string(encryptBigLetter(letter, rotate))
			continue
		}else {
			cipherText += string(letter)

		}
	}

	return cipherText
}



func encryptSmallLetter(r rune, shift int) rune {
	s := int(r) + (shift % 26)
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)

}

func encryptBigLetter(r rune, shift int) rune {
	s := int(r) + (shift % 26)
	if s > 'Z' {
		return rune(s - 26)
	} else if s < 'A' {
		return rune(s + 26)
	}
	return rune(s)

}

