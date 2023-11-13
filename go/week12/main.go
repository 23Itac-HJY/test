package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	as := a[:2]
	as[1] = "z"
	fmt.Println(a, as)

	b := [4]int{4, 3, 2, 1}
	bs := b[1:3]
	fmt.Println(b, bs)
}

// func main() {
// 	// var s []int
// 	// s = make([]int, 5)
// 	// s := make([]int, 5) // 단축연산자 사용, 변수 선언과 동시에 메모리 할당, 원소들은 제로값으로 초기화
// 	s := []int{0, 0, 0, 0, 0} //단축연산자 및 슬라이스 리터럴 사용, 변수 선언과 동시에 메모리 할당, 원소 초기화

// 	s[4] = 100
// 	s[0] = 7
// 	s[1] = 91

// 	for _, value := range s {
// 		fmt.Println(value)
// 	}

// 	copyS := s[1:4]
// 	for _, value := range copyS {
// 		fmt.Println(value)
// 	}

// 	test := [3]string{"inha", "go", "student"}
// 	testS := test[0:2] // testS := test[:2]
// 	testS2 := test[1:]

// 	//testS2[0] = "Python"
// 	//test[1] = "Python"
// 	testS[1] = "Python" // -2번등 음수 인덱스는 지원 X

// 	fmt.Println(test, len(test))
// 	fmt.Println(testS, len(testS))
// 	fmt.Println(testS2, len(testS2))
// }
