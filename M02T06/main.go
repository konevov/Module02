package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var t1 = time.Now()

func logTime() {
	t2 := time.Now()
	fmt.Println(" main function execution time: ", t2.Sub(t1))
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
