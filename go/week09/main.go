package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}
func main() {
	var a int = 5
	double(&a) // pass by pointer
	fmt.Printf("%d\n", a)
}

// package main

// import "fmt"

// func main() {
// 	a := 10
// 	var pa *int
// 	pa = &a

// 	fmt.Println(a, *pa)
// 	fmt.Println(&a, pa)
// 	fmt.Printf("%T %T %T %T\n", a, *pa, &a, pa)
// 	fmt.Println(&pa)

// }
