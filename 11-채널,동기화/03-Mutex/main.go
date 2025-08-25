package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // 뮤텍스 생성
	counter := 0

	// 1000개의 고루틴을 실행하여 counter를 1씩 증가시킴
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// counter에 접근하기 전에 Lock을 획득
			mu.Lock()
			counter++ // 이 연산은 원자적(atomic)이지 않음
			// Lock을 해제하여 다른 고루틴이 접근할 수 있도록 함
			mu.Unlock()
		}()
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 대기

	// mu.Lock()/Unlock()을 주석 처리하고 실행하면,
	// 데이터 경쟁으로 인해 counter 값이 1000이 아닐 수 있음.
	fmt.Printf("최종 counter 값: %d\n", counter)
}
