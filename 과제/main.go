package main

import "fmt"

func main() {
	var num int = 10
	var addr *int = &num

	*addr++

	fmt.Println("num:", num)
	fmt.Println("&num:", &num)
	fmt.Println("addr:", addr)
	fmt.Println("*addr:", *addr)
}
