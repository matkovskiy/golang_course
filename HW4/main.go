package main

import (
	"fmt"
	"hw4/internal/structs"
	"hw4/pkg/files_slices"
	"strconv"
	"strings"
)

func exercise1() {
	file := "assets/text.txt"
	search := "I'm addicted to you"
	lines := files_slices.Read_file(file)
	fmt.Printf("You are trying to find \"%s\"\n", search)
	result := files_slices.FindMatchInSlice(lines, search)
	if len(result) > 0 {
		fmt.Printf("We found %d matches\n", len(result))
		for i := range result {
			fmt.Printf("number #%d in line #%d\n", i, result[i]+1)
		}
	} else {
		fmt.Println("Sorry, we didn't find any matches")
	}
}

func exercise2() {
	subject := structs.Subject{}
	fmt.Println("Enter your Subject Name")
	fmt.Scanf("%s", &subject.Name)
	fmt.Println("Enter marks for %s. Type float(.), delimiter \",\"", subject.Name)
	var marks_raw string
	fmt.Scanf("%s", &marks_raw)

	for _, v := range strings.Split(marks_raw, ",") {
		value, _ := strconv.ParseFloat(v, 64)
		subject.Marks = append(subject.Marks, value)
	}
	fmt.Printf("You are selected subject: %s\n", subject.Name)
	var sum float64
	for i := range subject.Marks {
		sum = sum + subject.Marks[i]
	}
	average := sum / float64(len(subject.Marks))
	fmt.Printf("Average mark for %s is : %f\n", subject.Name, average)
}

func main() {
	fmt.Println("running exercise #1")
	exercise1()
	fmt.Println("\n\n\nRunning exercise #2")
	exercise2()

}
