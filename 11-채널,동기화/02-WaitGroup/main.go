package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup의 포인터와 작업자 ID를 인자로 받는 함수
func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	// 함수가 종료될 때 wg.Done()을 호출하여 작업 완료를 알림
	defer wg.Done()

	fmt.Printf("Worker %d: 작업 시작\n", id)
	time.Sleep(100 * time.Millisecond) // 작업 시뮬레이션
	fmt.Printf("Worker %d: 작업 완료\n", id)
}

func main() {
	// WaitGroup 생성
	var wg sync.WaitGroup

	// 3개의 작업자 고루틴을 실행
	for i := 1; i <= 3; i++ {
		// 기다려야 할 고루틴의 수를 1 증가시킴
		wg.Add(1)
		// worker 함수를 고루틴으로 실행
		go workerWithWaitGroup(i, &wg)
	}

	// 모든 고루틴이 Done()을 호출할 때까지 (카운터가 0이 될 때까지) 대기
	wg.Wait()

	fmt.Println("Main: 모든 작업이 완료되었습니다.")
}
