package main

import {
	"fmt"
}

func main() {
	var uni string = "inha"

	fmt.Println(uni)
	// var float64 float64 = 9.4
	// var test float64 = 7.9
	// fmt.Println(float64)
}


// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	//"math/rand"
// 	//"time"
// )

// // scanln()
// func main() {

// 	var number int

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number) // 앞에 것 = 개수

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0)
// 	}

// 	isPrime := true

// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // 첫 번째 약수가 발견되면 반복문 즉시 종료
// 		}
// 	}

// 	if isPrime { // 비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	//"math/rand"
// 	//"time"
// 	"bufio"
// 	"log"
// 	"os"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.6
// func main() {

// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	number, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // 첫 번째 약수가 발견되면 반복문 즉시 종료
// 		}
// 	}

// 	if isPrime { // 비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }
