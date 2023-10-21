package main

import (
	"HW11/pkg/files"
	"fmt"
	"regexp"
)

const numbers_files = "./assets/1689007675141_numbers.txt"
const text_files = "./assets/1689007676028_text.txt"

func exercise_1() {
	lines := files.Read_file(numbers_files)
	r, _ := regexp.Compile(`^\(?\d{3}\)?(\s|-|\s|\.?)\d{3}(-?|\.?|\s)\d{4}$`)
	for i := range lines {
		fmt.Printf("Cheking: %s. %t\n", lines[i], r.MatchString(lines[i]))
	}
}
func exercise_2() {
	lines := files.Read_file(text_files)
	r := regexp.MustCompile(`(ї|Ї)`)
	threesymbol := regexp.MustCompile(`(^|\s)[^\s|,]{3}(\s|,)`)

	var count int
	var countthreesymbol int
	for i := range lines {

		matches := r.FindAllStringSubmatch(lines[i], -1)
		matchesthreesymbol := threesymbol.FindAllStringSubmatch(lines[i], -1)
		count = count + len(matches)
		countthreesymbol = countthreesymbol + len(matchesthreesymbol)
	}
	fmt.Printf("We have %d of ї or Ї\n", count)
	fmt.Printf("We have %d of three symbols words", countthreesymbol)

}

func main() {
	fmt.Println("Running Exercise1")
	exercise_1()
	fmt.Println("\n========================\nRunning Exercise2")
	exercise_2()
}
