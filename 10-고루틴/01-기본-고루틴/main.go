package main

import (
	"fmt"
	time "time"
)

// "Hello" 메시지를 5번 출력하는 함수
func hello() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // 0.1초 대기
		fmt.Println("Hello from Goroutine!")
	}
}

func main() {
	// 1. 일반적인 함수 호출 (동기 방식)
	// hello() 함수가 모두 실행될 때까지 main 함수는 대기함
	// hello()

	// 2. 고루틴을 사용한 함수 호출 (비동기 방식)
	go hello() // hello() 함수를 새로운 고루틴에서 실행

	fmt.Println("Hello from main function!")

	// main 함수가 바로 종료되면, 생성된 고루틴이 실행될 기회조차 얻지 못하고 사라짐
	// 이를 방지하기 위해 임시로 1초간 대기 (좋은 방법은 아님)
	time.Sleep(1 * time.Second)

	fmt.Println("Main function finished")
}
