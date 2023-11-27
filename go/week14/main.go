package main

import "fmt"

func main() {
	balls := make(map[string]int)
	// var balls map[string]int
	fmt.Println(balls)
	fmt.Printf("%#v\n", balls)
	balls["성기훈"] = 20 // panic: assignment to entry in nil map <= 할당 안되면 에러
	fmt.Println(balls["성기훈"])
	fmt.Println(balls["오일남"])
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
