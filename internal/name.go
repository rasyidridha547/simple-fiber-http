package internal

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetName(name string) string {

	// since strings.Title deprecated in Go 1.18, I use cases
	caser := cases.Title(language.AmericanEnglish)

	str := caser.String(name)

	result := fmt.Sprintf("Hello, %+v", str)

	return result
}
