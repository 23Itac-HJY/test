package main

import (
	"fmt"
)

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	// var primes [3]int = [3]int{2, 3, 5} //배열 리터럴로 배열 초기화

	primes := [3]int{2, 3, 5}

	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true} // 초기화하지 않은 배열 원소 나머지 두 값에 제로값(false) 할당
	fmt.Println(test[3])

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	// fmt.Println(primes[3]) // invalid argument: index 3 out of bounds [0:3]

	i := 0
	for i < 4 { //panic: runtime error: index out of range [3] with length 3
		fmt.Println(primes[i])
		i++
	}

}
