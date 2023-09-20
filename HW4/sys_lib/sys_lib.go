package sys_lib

import (
	"encoding/json"
	"fmt"
	"hw4/structs"
	"io"
	"os"
)

func Init_input() (result structs.Student) {
	new := structs.Student{}
INIT_INPUT:
	fmt.Println("Enter your Name and Surname")
	fmt.Scanf("%s %s", &new.Name, &new.Surname)
	fmt.Printf("You are %s %s.\n Is it correct? (1: Yes, any other chart - false)\n", new.Surname, new.Name)

	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "1" {
		fmt.Println("Moving forward....\n")
	} else {
		fmt.Println("anwser is no.\n Let's try again")
		goto INIT_INPUT
	}
	result = new
	return
}

func Read_subjects(filename string) (result structs.Subjects) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)

	defer jsonFile.Close()
	return

}

func Print_stundet_subjects(student *structs.Student) {
	for i, v := range student.Subjects {
		fmt.Printf("Your subject #%d is %s\n", i, v.Name)

	}
}

func Who_you_are(student *structs.Student) {
	tech := 0
	for _, v := range student.Subjects {
		if v.Techical {
			tech++
		}
	}
	if tech > 2 {
		fmt.Println("Congratulate you are enginer")
	} else {
		if tech == 2 {
			fmt.Printf("Congratulate you are TechLead")
		} else {
			fmt.Printf("Sorry, buy you are manager")
		}
	}
}
