package main

import (
	"fmt"
	"strings"
	//reflect
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	// // zero value, 초기화를 안하면 갖는 값
	// var c rune // = '가'
	// var a int  // 7
	// var b float64
	// var d bool
	// var e string

	// // fmt.Printf("%c\n", c)
	// // fmt.Printf("%T\n", c)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))

}
