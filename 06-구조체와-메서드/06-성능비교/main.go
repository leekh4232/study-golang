package main

import (
	"fmt"
	"time"
)

// 큰 구조체 예제
type LargeStruct struct {
	Data [1000]int
	Name string
	ID   int
}

// 값 수신자 메서드
func (ls LargeStruct) ProcessWithValue() int {
	sum := 0
	for _, v := range ls.Data {
		sum += v
	}
	return sum
}

// 포인터 수신자 메서드
func (ls *LargeStruct) ProcessWithPointer() int {
	sum := 0
	for _, v := range ls.Data {
		sum += v
	}
	return sum
}

func main() {
	// 큰 구조체 인스턴스 생성
	large := LargeStruct{
		Name: "TestStruct",
		ID:   1,
	}

	// 배열에 데이터 채우기
	for i := 0; i < 1000; i++ {
		large.Data[i] = i
	}

	iterations := 100000

	// 값 수신자 성능 측정
	start := time.Now()
	for i := 0; i < iterations; i++ {
		large.ProcessWithValue()
	}
	valueTime := time.Since(start)

	// 포인터 수신자 성능 측정
	start = time.Now()
	for i := 0; i < iterations; i++ {
		large.ProcessWithPointer()
	}
	pointerTime := time.Since(start)

	fmt.Printf("=== 성능 비교 결과 (%d회 반복) ===\n", iterations)
	fmt.Printf("값 수신자 시간: %v\n", valueTime)
	fmt.Printf("포인터 수신자 시간: %v\n", pointerTime)
	fmt.Printf("성능 차이: %.2fx\n", float64(valueTime)/float64(pointerTime))
}
