package patterns

import "strings"

var bunner string = "\n\n\n==================\nThe place for your advertisment\n\n\n==================\n"

func AddBaner(text []string) []string {
	return append(text, bunner)

}

func LowerCase(text []string) []string {
	var output []string
	for i := range text {
		output = append(output, strings.ToLower(text[i]))
	}
	return output
}
