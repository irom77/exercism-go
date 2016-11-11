package acronym

import (
	"strings"
	//"fmt"
	//"unicode"
	"unicode"
)

const targetTestVersion = 1
// fails in the case below but I think expected should be PHPHP
//acronym_test.go:33: Acronym test [PHP: Hypertext Preprocessor], expected [PHP], actual [PHPHP]


func Acronym(s string) string{
var result string
	//words := strings.Split(s, " ")
	// Return true if whitespace or - char.
	f := func(c rune) bool {
		return c == ' ' || c == '-'
	}
	words := strings.FieldsFunc(s, f)
	//fmt.Println(words)
	for _, word := range words {
		for i, letter := range word {
			if i==0 {
				result += string(letter)
				//break
			} else if unicode.IsUpper(letter) {
				result += string(letter)
			}
		}
	}
	return strings.ToUpper(result)
}
