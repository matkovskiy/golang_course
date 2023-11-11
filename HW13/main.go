package main

import (
	files_organizer "HW13/pkg/files_organizer"
	"fmt"
	"strings"
)

var path string
var answer, input string

func interactiveWithUser() {
INIT_INPUT:
	fmt.Println(`Enter action.\n
		1 - Create File,
		2 - List All files, directories in folder,
		3 - List with pattern files, directories in folder,
		4 - Delete all files in directories,
		5 - Delete files in directories by pattern`)
	fmt.Scanf("%s", &input)

	fmt.Println(`Enter path:`)
	fmt.Scanf("%s", &path)

	fmt.Println(`Is everything correct?`)
	fmt.Scanf("%s", &answer)
	if answer == "1" {
		fmt.Println("Moving forward....")
	} else {
		fmt.Println("answer is no.\n Let's try again")
		goto INIT_INPUT
	}

	switch input {
	case "1":
		var files string
		fmt.Println(`Enter file (delimiter is ","):`)
		fmt.Scanf("%s", &files)
		files_organizer.CreateFiles(path, strings.Split(files, ","))

	case "2":
		output := files_organizer.ScanDirWithPattern(path, `.*`)
		for i := range output {
			fmt.Println(output[i])
		}

	case "3":
		var pattern string
		fmt.Println(`Enter pattern:`)
		fmt.Scanf("%s", &pattern)
		output := files_organizer.ScanDirWithPattern(path, pattern)

		for i := range output {
			fmt.Println(output[i])
		}
	case "4":
		files_organizer.DeleteFilesWithPattern(path, `.*`)

	case "5":
		var pattern string
		fmt.Println(`Enter pattern:`)
		fmt.Scanf("%s", &pattern)
		files_organizer.DeleteFilesWithPattern(path, `.*`)
	}
}

func main() {
	interactiveWithUser()
}
