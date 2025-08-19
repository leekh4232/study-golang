package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	b := a
	fmt.Println(b)
	// 배열의 값을 변경해도 복사본에는 영향을 주지 않음
	b[0] = 10
	fmt.Println(a)
	fmt.Println(b)
}
