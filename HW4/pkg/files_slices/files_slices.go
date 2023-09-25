package files_slices

import (
	"bufio"
	"fmt"
	"os"
)

func Read_file(file string) (lines []string) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return
}

func FindMatchInSlice(input_slice []string, search_str string) (output []int) {

	for i, v := range input_slice {
		if v == search_str {
			output = append(output, i)
		}
	}
	return
}
