package files

import (
	"bufio"
	"fmt"
	"os"
)

type Text interface {
	DisplayText()
}

type GeneralText struct {
	Data *[]string
}

func (text GeneralText) Display() {
	fmt.Println(text.Data)
}

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
