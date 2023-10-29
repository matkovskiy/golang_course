package patterns

import (
	"strings"
)

type TextManipulator func([]string) []string

type TextManipulatorDecorator func(TextManipulator) TextManipulator

func Ident(data []string) []string {
	return data
}
func ToLower(m TextManipulator) TextManipulator {
	return func(s []string) []string {
		var output []string
		for i := range s {
			output = append(output, strings.ToLower(s[i]))
		}
		return output
		// output
	}
}

func AppendDecoratorFactory(banner string) TextManipulatorDecorator {
	return func(m TextManipulator) TextManipulator {
		return func(s []string) []string {
			return append([]string{banner}, s...)
		}
	}
}
