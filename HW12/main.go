package main

import (
	"HW12/pkg/files"
	"HW12/pkg/patterns"
	"bufio"
	"fmt"
	"os"
)

const text_files = "./assets/1689007676028_text.txt"

type Answers struct {
	first, second int
}

func interactiveWithUser() {
	ans := Answers{}
	body := files.Read_file(text_files)
	var text []string
	fmt.Println("Please choose your decorator. 1 - Add banner, 2 - replace text to lowercase")
	fmt.Scanf("%d", &ans.first)
	switch ans.first {
	case 1:
		text = patterns.AddBaner(body)
	case 2:
		text = patterns.LowerCase(body)
	default:
		panic("Unexpected input")
	}
	fmt.Println("Printing new text:")
	for i := range text {
		fmt.Println(text[i])
	}
	fmt.Print("\n\n\n")

	fmt.Println("Please choose your modification. 1 - Duplicate spaces, 2 - Remove specific word")
	fmt.Scanf("%d", &ans.second)
	switch ans.second {
	case 1:
		myText := patterns.Texting{Text: patterns.TextWithDoubleSpaces{}}
		text = myText.Text.Processing(text, "null")

	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter word for removing: ")
		word, _ := reader.ReadString('\n')
		fmt.Printf("Word is:%s", word)
		myText := patterns.Texting{Text: patterns.RemoveWord{}}
		text = myText.Text.Processing(text, word[:len(word)-1])

	default:
		panic("Unexpected input")
	}
	fmt.Println("Priting final result:")
	for i := range text {
		fmt.Println(text[i])
	}
}

func main() {
	interactiveWithUser()
}
