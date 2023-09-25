package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n') // option 1 에러 무시하고 진행하라는 의미
	fmt.Println(inputScore)

	// var now time.Time
	// now = time.Now()
	// var year = now.Year()
	// month := now.Month()
	// fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
