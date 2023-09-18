package main

import (
	"fmt"
	"hw3/sys_lib"
	"time"
)

func main() {

	filename := "assets/hw3.json"
	student := sys_lib.Init_input()
	fmt.Printf("Wellcome, %s\n", student.Name)
	subjects := sys_lib.Read_subjects(filename)

Make_choise:
	student.Subjects = nil

	fmt.Printf("List of subjects available to choose\n===================\n")
	for i := 0; i < len(subjects.Subjects); i++ {
		fmt.Printf("#%d Subject name: %s, Techical: %t, Complexity: %d\n", i, subjects.Subjects[i].Name, subjects.Subjects[i].Techical, subjects.Subjects[i].Complexity)
	}
	var choise int
	for j := 1; j < 5; j++ {
		fmt.Printf("Welcome to level#%d\nSelect subject for leveling: (0-6)\n", j)
		fmt.Scanf("%d", &choise)
		student.Learn_subject(subjects.Subjects[choise])
	}
	// student.Learn_subject(subjects.Subjects[0])
	sys_lib.Print_stundet_subjects(&student)

	var answer string
	fmt.Println("Do you like your skillset ? (1: Yes, any other character - no)")
	fmt.Scanf("%s", &answer)
	if answer == "1" {
		fmt.Println("Moving forward....")
	} else {
		fmt.Println("Your anwser is no.\n Let's make a choice again.")
		time.Sleep(3 * time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Scanf("%d")
		goto Make_choise
	}
	sys_lib.Who_you_are(&student)

}
