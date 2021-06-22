package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		panic(err)
	}
	defer fileOpen.Close()

	fileCreate, err := os.Create("data/data_out.txt")
	if err != nil {
		panic(err)
	}
	defer fileCreate.Close()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			data, err := ioutil.ReadFile("data/data_out.txt")
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			fmt.Println("Contents of file:", string(data))
		}

	}()

	lineCount := 0
	str := bufio.NewScanner(fileOpen)

	for str.Scan() {
		str := str.Text()
		sliceStr := strings.Split(str, "|")
		name, address, city := sliceStr[0], sliceStr[1], sliceStr[2]
		lineCount++

		if name == "" || address == "" || city == "" {
			err := fmt.Sprintf("parse error: empty field on string %d", lineCount)
			panic(err)
		}

		newStr := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", lineCount, name, address, city)

		fmt.Fprintln(fileCreate, newStr)
	}
}
