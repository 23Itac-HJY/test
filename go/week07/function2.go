package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {

	if n < 2 {
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil // 소수 판정 값, 정상데이터
}

// 소수 판정 프로그램 v1.4, isPrime 함수 안헤 변수를 하나 줄이고, return 구문 추가, break제거
func main() {
	var a, b int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil {
		log.Fatal(err)
	}

	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, ", ")
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true

// 	if n < 2 {
// 		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}

// 	return prime, nil // 소수 판정 값, 정상데이터
// }

// // 소수 판정 프로그램 v1.1 : 함수 적용, error 리턴
// func main() {
// 	var a, b int

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if a>b{
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, ", ")
// 		}
// 	}
// }
