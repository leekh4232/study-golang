package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // 3초 걸리는 작업 --> 타임아웃이 먼저 발생하여 실행되지 않음
		fmt.Println("작업 완료")
	case <-ctx.Done(): // 타임아웃 또는 취소
		fmt.Println("작업 중단:", ctx.Err())
	}
}

func main() {
	// 타임아웃 2초짜리 컨텍스트
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	task(ctx)
}
