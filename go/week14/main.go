package main

import "fmt"

func status(name string) {
	balls := map[string]int{
		"성기훈": 20,
		"오일남": 0,
	}
	//ball := balls[name]
	ball, exists := balls[name] // exists 는 써도 되고 안써도 됨, bool 타입
	//fmt.Println(ball, exists)
	if !exists {
		fmt.Println(name, "님은 게임에 참여하실 수 없습니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은 ", balls[name], "개로 탈락")
	} else {
		fmt.Println(name, "님은 게임에서 승리하였습니다.")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("성기훈")
	// balls := make(map[string]int)
	// // var balls map[string]int
	// fmt.Println(balls)
	// fmt.Printf("%#v\n", balls)
	// balls["성기훈"] = 20 // panic: assignment to entry in nil map <= 할당 안되면 에러
	// fmt.Println(balls["성기훈"])
	// fmt.Println(balls["오일남"])
}

// package main

// import "fmt"

// func main() {
// 	// var games map[int]string
// 	// games = make(map[int]string)  1st
// 	//games := make(map[int]string)  2nd
// 	games := map[int]string{
// 		456: "성기훈",
// 		218: "박해수",
// 		067: "강새벽"}
// 	//append

// 	games[456] = "성기훈"
// 	games[218] = "박해수"
// 	games[067] = "강새벽"
// 	games[001] = "오일남"
// 	games[199] = "알리"
// 	games[101] = "lee"

// 	// fmt.Println(games[067])
// 	for _, v := range games {
// 		fmt.Println(v)
// 	}

// 	// update
// 	games[101] = "장덕수"
// 	for k, v := range games {
// 		fmt.Println(k, v)
// 	}

// 	//delete
// 	delete(games, 199)
// 	for k, v := range games {
// 		fmt.Println(k, v)
// 	}
// }
