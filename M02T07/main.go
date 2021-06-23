package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := fileOpen.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileCreate, err := os.Create("data/data_out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := fileCreate.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			fmt.Println(panicErr)
			data, err := ioutil.ReadFile("data/data_out.txt")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(data))
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
			err := fmt.Sprintf("parse error: empty field on string %d\n", lineCount)
			panic(err)
		}

		newStr := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", lineCount, name, address, city)

		fmt.Fprintln(fileCreate, newStr)
	}
}
