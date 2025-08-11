package main

import (
	"fmt"
	"runtime"
	"time"
)

// 대량의 문자열 슬라이스를 생성하여 메모리 사용량을 늘리는 함수
func allocateMemory() {
	var s []string
	for i := 0; i < 1000000; i++ { // 100만 개의 문자열 생성
		s = append(s, fmt.Sprintf("Hello, Go GC! %d", i))
	}
	// 함수 종료 후 s는 더 이상 참조되지 않으므로 GC 대상이 된다.
	fmt.Println("  메모리 할당 완료. (GC 대상)")
}

func main() {
	fmt.Println("---", "6.1 GC 동작 확인 (간접적)", "---")

	// 초기 메모리 상태 확인
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("초기 메모리 사용량 (HeapAlloc): %v MB\n", m.HeapAlloc/1024/1024)

	// 대량의 메모리 할당 함수 호출
	allocateMemory()

	// GC가 동작할 시간을 주기 위해 잠시 대기
	time.Sleep(2 * time.Second)

	// GC 호출 (명시적 호출은 권장되지 않지만, 테스트 목적으로 사용)
	runtime.GC()
	fmt.Println("  GC 강제 호출 완료.")

	// GC 후 메모리 상태 확인
	runtime.ReadMemStats(&m)
	fmt.Printf("GC 후 메모리 사용량 (HeapAlloc): %v MB\n", m.HeapAlloc/1024/1024)

	fmt.Println("\n---", "6.2 GC의 자동 동작", "---")
	fmt.Println("Go의 GC는 대부분 자동으로 백그라운드에서 동작하며, 개발자가 명시적으로 호출할 필요는 없다.")
	fmt.Println("메모리 사용량이 일정 수준을 넘거나, 특정 조건이 되면 자동으로 실행된다.")
}
