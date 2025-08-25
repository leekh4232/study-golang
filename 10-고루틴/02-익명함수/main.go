package main

import (
	"fmt"
	time "time"
)

func main() {
	// 잘못된 예: 반복 변수 i를 직접 사용
	fmt.Println("잘못된 경우:")
	for i := 0; i < 3; i++ {
		// 이 고루틴은 i의 메모리 주소를 참조함
		// 고루틴이 실행될 시점에는 반복문이 끝나 i는 3이 되어 있을 가능성이 높음
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

	fmt.Println("\n올바른 예 (1): 변수 복사")
	// 올바른 예 (1): 반복 변수를 고루틴 내의 지역 변수로 복사
	for i := 0; i < 3; i++ {
		i := i // 새로운 변수 i에 현재 i의 값을 복사 (Shadowing)
		go func() {
			fmt.Println(i) // 복사된 변수를 사용
		}()
	}
	time.Sleep(1 * time.Second)

	fmt.Println("\n올바른 예 (2): 인자로 전달")
	// 올바른 예 (2): 반복 변수를 고루틴 함수의 인자로 전달
	for i := 0; i < 3; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i) // 현재 i의 값을 인자로 전달
	}
	time.Sleep(1 * time.Second)
}
