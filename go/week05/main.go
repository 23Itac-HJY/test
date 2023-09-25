package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess number game!")
	answer := rand.Intn(100) + 1 // 1~100
	fmt.Println(answer)

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

}
