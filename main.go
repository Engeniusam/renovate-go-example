package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	tag, err := language.Parse("en-US")
	if err != nil {
		fmt.Println("Error parsing language tag:", err)
	} else {
		fmt.Println("Parsed language tag:", tag)
	}
}
