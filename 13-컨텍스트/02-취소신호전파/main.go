package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 취소 신호 수신
			fmt.Printf("[%s] 작업 중단: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] 작업 진행 중...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// 1. 최상위 컨텍스트 생성
	ctx := context.Background()

	// 2. 취소 가능한 컨텍스트 생성
	ctx, cancel := context.WithCancel(ctx)

	// 3. 두 개의 worker 고루틴 실행
	go worker(ctx, "A")
	go worker(ctx, "B")

	// 4. 2초 뒤 작업 취소
	time.Sleep(2 * time.Second)
	fmt.Println("메인: 취소 신호 전송")
	cancel()

	// 5. 고루틴 종료 대기
	time.Sleep(1 * time.Second)
}
