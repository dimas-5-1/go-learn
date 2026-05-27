package main

import (
	"fmt"
)

// checks doubled ASCII letters in string (case insensitive)

func have_duplicated_letters(str string) bool {

	if len(str) == 1 {
		return false
	}

	for i := 1; i < len(str); i++ {
		if int(str[0])&^32 == int(str[i])&^32 {
			return true
		} else {
			have_duplicated_letters(str[1:])
		}
	}

	return false
}

func main() {

	str := "Golang"

	fmt.Println("String ", str, ", duplicated letters: ", have_duplicated_letters(str))

}
