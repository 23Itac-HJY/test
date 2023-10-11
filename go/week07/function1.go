package main

import "fmt"

func processScore(kor int, eng int, math int) (int, int) {
	totalScore := kor + eng + math
	average := totalScore / 3

	return totalScore, average
}

func main() {
	t, a := processScore(100, 90, 93)
	fmt.Printf("%s의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길동", t, a)
	t, a = processScore(89, 91, 92)
	fmt.Printf("%s의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길순", t, a)
}

// package main

// import "fmt"

// func processScore(name string, kor int, eng int, math int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	fmt.Printf("%s의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	processScore("홍길동", 100, 90, 93)
// 	processScore("홍길순", 89, 91, 92)
// }

// package main

// import "fmt"

// func sayHello() {
// 	fmt.Println("Hello World!")
// }

// func main() {
// 	sayHello()
// }
