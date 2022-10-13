package DB

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Read the txt file and set the value into the database.
func ReadTxt() {
	file, err := os.Open("es.txt")
	if err != nil {
		log.Fatal("The name file is not the correct one")
	}
	//* Close the file
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//TODO Change this to handle the insertion of the value for the redis database.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
