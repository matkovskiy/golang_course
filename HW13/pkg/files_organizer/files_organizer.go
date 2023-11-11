package files_organizer

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func CreateFiles(path string, files []string) {
	for i := range files {
		file, err := os.Create(path + files[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("File %s%s is created successfully.\n", path, files[i])
	}

}

func DeleteFiles(path string, files []string) {
	for i := range files {
		err := os.Remove(path + files[i])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("File %s%s is deleted successfully.\n", path, files[i])
	}
}

func ScanDirWithPattern(path, pattern string) []string {
	r, _ := regexp.Compile(pattern)
	var output []string

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		if r.MatchString(e.Name()) {
			output = append(output, e.Name())
		}
	}
	return output
}

func DeleteFilesWithPattern(path, pattern string) {
	input := ScanDirWithPattern(path, pattern)
	for i := range input {
		err := os.Remove(path + input[i])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("File %s%s is deleted successfully.\n", path, input[i])
	}
}
