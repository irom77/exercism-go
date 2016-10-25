package acronym

import (
	"strings"
	"fmt"
)

const targetTestVersion = 1

func Acronym(s string) string{
var result string
	words := strings.Split(s, "")
	for _, word := range words {
		result = fmt.Sprint(word[0])
	}
	return result
}
