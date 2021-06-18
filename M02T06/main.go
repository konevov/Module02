package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func main() {
	defer duration(track(" main function execution time: "))

	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		panic(err)
	}
	defer fileOpen.Close()

	fileCreate, err := os.Create("data/out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lineCount := 0
	defer func() {
		stat, err := fileCreate.Stat()
		if err != nil {
		}
		fmt.Printf("The file is %d bytes long\n", stat.Size())
		fmt.Printf("The number of lines in the file %d\n", lineCount)
		fileCreate.Close()
	}()

	str := bufio.NewScanner(fileOpen)
	for str.Scan() {
		lineCount++
		newStr := fmt.Sprintf("%d. %s", lineCount, str.Text())
		fmt.Fprintln(fileCreate, newStr)
	}

}
