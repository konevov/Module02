package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fileOpen.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(fileOpen)
	_, isPrefix, err := r.ReadLine()
	counter := 0

	for err == nil && !isPrefix {
		counter++
		_, isPrefix, err = r.ReadLine()
	}

	if isPrefix {
		fmt.Printf("Buffer size too small in string %d", counter)
		return
	}

	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total strings: %d", counter)

}
