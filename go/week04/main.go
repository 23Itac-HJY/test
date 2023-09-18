package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// var a int // declaration
	// a = 7     // assign value
	// var a int = 7 // declaration & assign value
	// var a = 7
	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	// b := 8.34 // float64
	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)

	var c7 string // 변수명은 알파벳 대소문자로 시작해야 한다.
	var d int
	var e bool
	var f float64
	var G = 99           // 대문자로 시작한 G만 외부에서 접근 가능, 함수도 마찬가지
	koreaZombie := "정찬성" //koreazombie := "정찬성" 문법적으로 문제 없지만, camel케이스를 사용하는 관례를 따르자
	fmt.Println(koreaZombie)

	fmt.Println(c7, d, e, f, G) // 제로값 확인

	fmt.Println('Z', '2', '\n', '김', '인', '하') //rune literals(int32) 90 50 10 44,608 51,064 54,616
	fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf(2), reflect.TypeOf("HI"), reflect.TypeOf(4.99), reflect.TypeOf(false))
	fmt.Println(math.Floor(2.75), math.Ceil(2.75), math.Sqrt(16))
	fmt.Println(strings.Title("\\open source\tprogramming\n\"go\"!"))
}
