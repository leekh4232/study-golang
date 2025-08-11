package main

import "fmt"

// 이 함수는 int 포인터를 반환하므로, x는 힙에 할당될 가능성이 높다 (이스케이프).
func createIntPointer() *int {
	x := 10 // x는 함수 내부에서 선언되었지만, 주소가 반환되므로 힙으로 이스케이프될 수 있다.
	fmt.Println("  createIntPointer 내부 - x의 주소:", &x)
	return &x
}

// 이 함수는 int 값을 반환하므로, y는 스택에 할당될 가능성이 높다.
func createIntValue() int {
	y := 20                                               // y는 함수 종료 후 더 이상 참조되지 않으므로 스택에 할당될 가능성이 높다.
	fmt.Println("  createIntValue 내부 - y의 주소 (개념적):", &y) // 실제 주소는 스택 프레임 내
	return y
}

// 큰 구조체를 값으로 반환하는 경우 (복사 비용 발생)
type LargeStruct struct {
	Data [1024]byte // 1KB 크기의 배열
}

func createLargeStructByValue() LargeStruct {
	var s LargeStruct
	// s에 데이터 채우기
	return s // s 전체가 복사되어 반환된다.
}

// 큰 구조체를 포인터로 반환하는 경우 (복사 비용 감소, 힙 할당 가능성 높음)
func createLargeStructByPointer() *LargeStruct {
	s := new(LargeStruct) // 힙에 할당될 가능성 높음
	// s에 데이터 채우기
	return s // 포인터 값만 복사되어 반환된다.
}

func main() {
	// 5.1 이스케이프 분석 예시
	fmt.Println("---", "5.1 이스케이프 분석 예시", "---")
	ptr := createIntPointer()
	fmt.Println("main 함수 - ptr이 가리키는 값:", *ptr) // 10
	*ptr = 15
	fmt.Println("main 함수 - ptr이 가리키는 값 변경 후:", *ptr) // 15

	val := createIntValue()
	fmt.Println("main 함수 - val:", val) // 20

	// 5.2 큰 데이터 구조 전달 방식에 따른 메모리 효율성 (개념적)
	fmt.Println("\n---", "5.2 큰 데이터 구조 전달 방식에 따른 메모리 효율성", "---")
	// 값으로 전달: LargeStruct 전체가 복사된다. (비용 높음)
	largeVal := createLargeStructByValue()
	fmt.Printf("LargeStruct (값): %p\n", &largeVal)

	// 포인터로 전달: 포인터 값만 복사된다. (비용 낮음)
	largePtr := createLargeStructByPointer()
	fmt.Printf("LargeStruct (포인터): %p\n", largePtr)
}
