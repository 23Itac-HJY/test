package main

import (
	"bufio"
	"fmt"
	"log" //Fatal 함수
	"os"
	"strconv" // ParseInt
	"strings" // TrimSpace
)

func main() {
	fmt.Print("구구단 단 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n') // in은 현재 string

	if err != nil { // 조건문에 () 사용 안함
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)
	dan, err := strconv.ParseInt(in, 10, 32) // 32bit로 했지만 int64 형으로 리턴
	// dan, err := strconv.Atoi(in)  // 위와 동일, atoi는 int로 변환
	if err != nil {
		log.Fatal(err)
	}
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(dan, " * ", i, " = ", (int(dan) * i))
	// }
	i := 1
	for i < 10 { // 다른 언어의 while과 동일
		fmt.Println(dan, " * ", i, " = ", (int(dan) * i))
		i++
	}

}
