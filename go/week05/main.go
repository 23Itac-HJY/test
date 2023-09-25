package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n') //assignment mismatch: 1 variable but reader.ReadString returns 2 values
	fmt.Println(inputScore)

	// var now time.Time
	// now = time.Now()
	// var year = now.Year()
	// month := now.Month()
	// fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
