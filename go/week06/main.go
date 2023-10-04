package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가'      // 작은 따옴표 => 룬, 큰따옴표 => 문자열
	fmt.Println(c)        //유니코드값(UTF-8) 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune = int32 alias int32의 별명

	fmt.Println(math.Ceil(2.71))
	fmt.Println(math.Floor(2.71))
	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github"))

}
