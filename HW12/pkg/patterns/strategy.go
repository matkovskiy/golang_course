package patterns

import "strings"

type Text interface {
	Processing([]string, string) []string
}

type Texting struct {
	Text Text
}

type TextWithDoubleSpaces struct{}

func (TextWithDoubleSpaces) Processing(text []string, word string) []string {
	var output []string
	for i := range text {
		output = append(output, strings.Replace(text[i], " ", "  ", -1))
	}
	return output
}

type RemoveWord struct{}

func (RemoveWord) Processing(text []string, word string) []string {
	var output []string
	for i := range text {
		output = append(output, strings.ReplaceAll(text[i], word, ""))
	}
	return output
}
