package mymath

func MyAbs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func MyPower(base int, exponent int) int { // 외부에 함수를 공개하려면 반드시 함수 이름을 대문자로 시작
	result := 1
	for i := 0; i < exponent; i++ {
		result = result * base
	}
	return result
}
