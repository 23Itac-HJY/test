package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      // remove space
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // string to 32bit float
	if err != nil {
		log.Fatal(err)
	}
	var grade string
	if inputScore >= 90 { // invalid operation: inputScore >= 90 (mismatched types string and untyped int)
		grade = "A Grade!"
		// fmt.Println("You got", grade)
	} else {
		grade = "under A Grade..."
		// fmt.Println("You got", grade)
	}
	fmt.Println(grade)

	// var now time.Time
	// now = time.Now()
	// var year = now.Year()
	// month := now.Month()
	// fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
