package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램
func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0

	number := rand.Intn(150) + 2 // 0과 1 제외, 2~151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}

}
