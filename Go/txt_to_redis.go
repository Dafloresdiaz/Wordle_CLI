package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readTxt() {
	file, err := os.Open("es.txt")
	if err != nil {
		log.Fatal("The name file is not the correct one")
	}
	//* Close the file
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	readTxt()
}
