package main

import "fmt"

func swap(n1 *int, n2 *int) { // 두 수를 교환
	temp := *n1
	*n1 = *n2
	*n2 = temp

}
func main() {
	var a int = 10
	b := 20
	c := &a // var c *int = &a

	//fmt.Printf("%d %d %x\n", a, b, c)
	fmt.Printf("%d %d %d\n", a, b, *c)
}

// package main

// import "fmt"

// func double(n *int) {
// 	*n = *n * 2
// }
// func main() {
// 	var a int = 5
// 	double(&a) // pass by pointer
// 	fmt.Printf("%d\n", a)
// }

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
