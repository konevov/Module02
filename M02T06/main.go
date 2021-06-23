package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var tStart = time.Now()

func logTime() {
	tEnd := time.Now()
	fmt.Println("Main function execution time: ", tEnd.Sub(tStart))
}

func main() {
	defer logTime()

	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		fmt.Println(err)
		return
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
			fmt.Println(err)
			return
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
