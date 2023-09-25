package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // option 2
	log.Fatal(err)                             // 프로그램 종료해서 점수 출력 안됨
	fmt.Println(inputScore)

	// var now time.Time
	// now = time.Now()
	// var year = now.Year()
	// month := now.Month()
	// fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
