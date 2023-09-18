package sys_lib

import (
	"fmt"
	"hw3/structs"
)

func Init_input() (result structs.Engineer) {
	var new = structs.Engineer{}
INIT_INPUT:
	fmt.Println("Enter your name, age and Experience ")
	fmt.Scanf("%s %d %d", &new.Name, &new.Age, &new.Experience)
	fmt.Printf("You are %s,  %d years old with %d year of experience\n", new.Name, new.Age, new.Experience)

	var answer string

	fmt.Println("Is it correct? (1: Yes, any other chart - false) ")
	fmt.Scanf("%s", &answer)
	if answer == "1" {
		fmt.Println("anwser is yes")
	} else {
		fmt.Println("anwser is no.\n Let's try again")
		goto INIT_INPUT
	}
	result = new
	return
}

func Ski() (result structs.Engineer) {
	var new = structs.Engineer{}
INIT_INPUT:
	fmt.Println("Enter your name, age and Experience ")
	fmt.Scanf("%s %d %d", &new.Name, &new.Age, &new.Experience)
	fmt.Printf("You are %s,  %d years old with %d year of experience\n", new.Name, new.Age, new.Experience)

	var answer string

	fmt.Println("Is it correct? (1: Yes, any other chart - false) ")
	fmt.Scanf("%s", &answer)
	if answer == "1" {
		fmt.Println("anwser is yes")
	} else {
		fmt.Println("anwser is no.\n Let's try again")
		goto INIT_INPUT
	}
	result = new
	return
}
