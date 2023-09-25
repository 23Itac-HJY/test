package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // err declared and not used
	fmt.Println(inputScore)

	// var now time.Time
	// now = time.Now()
	// var year = now.Year()
	// month := now.Month()
	// fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
