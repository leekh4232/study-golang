package main

import (
	"context"
	"fmt"
)

// 1. 키 충돌을 방지하기 위해 사용자 정의 타입을 사용
type requestIDKey string

func processRequest(ctx context.Context) {
	// 2. 컨텍스트에서 "requestID" 키로 값을 조회
	id, ok := ctx.Value(requestIDKey("requestID")).(string)
	if !ok {
		id = "unknown"
	}
	fmt.Printf("Processing request with ID: %s\n", id)
}

func main() {
	// 3. 비어 있는 최상위 컨텍스트 생성
	ctx := context.Background()

	// 4. 컨텍스트에 "requestID"와 값 "12345"를 저장
	ctxWithID := context.WithValue(ctx, requestIDKey("requestID"), "12345")

	// 5. 값이 저장된 컨텍스트를 함수에 전달
	processRequest(ctxWithID)
}
