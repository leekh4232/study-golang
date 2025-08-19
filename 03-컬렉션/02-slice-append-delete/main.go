package main

import "fmt"

func main() {
	/** 슬라이스에 요소 추가 (append) */

	// `append` 함수는 슬라이스에 요소를 추가하고, 새로운 슬라이스를 반환한다.
	// 따라서 반환 값을 반드시 기존 슬라이스 변수에 재할당해야 한다.
	// 용량이 부족하면 Go 런타임이 자동으로 더 큰 내부 배열을 할당하고 기존 요소를 복사한다.
	// Java의 `ArrayList.add()`는 내부적으로 배열 크기를 늘리지만, Go는 명시적으로 재할당해야 한다.
	slice := []int{1, 2, 3}
	fmt.Printf("초기 슬라이스: %v, 길이: %d, 용량: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 4) // 단일 요소 추가. Java의 `list.add(4);`와 유사.
	fmt.Printf("4 추가 후: %v, 길이: %d, 용량: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 5, 6) // 여러 요소 추가. `append`는 가변 인자(variadic arguments)를 받는다.
	fmt.Printf("5, 6 추가 후: %v, 길이: %d, 용량: %d\n", slice, len(slice), cap(slice))

	anotherSlice := []int{7, 8}
	// `...` 연산자를 사용하여 다른 슬라이스의 모든 요소를 현재 슬라이스에 추가한다.
	// Java의 `list.addAll(anotherList);`와 유사하다.
	slice = append(slice, anotherSlice...)
	fmt.Printf("다른 슬라이스 추가 후: %v, 길이: %d, 용량: %d\n", slice, len(slice), cap(slice))

	/** 슬라이스에서 요소 삭제 */

	// Go는 Java의 `list.remove(index)`나 `list.remove(Object)`와 같은 직접적인 삭제 메서드를 제공하지 않는다.
	// 슬라이싱(slicing)과 `append`를 조합하여 요소를 삭제하는 것이 일반적인 방법이다.

	// 인덱스 1의 요소 삭제 (값 20 삭제)
	sliceToDelete := []int{10, 20, 30, 40, 50}
	fmt.Println("삭제 전 슬라이스:", sliceToDelete)
	// `sliceToDelete[:1]`은 인덱스 0까지의 슬라이스([10])를 생성한다.
	// `sliceToDelete[2:]`는 인덱스 2부터 끝까지의 슬라이스([30, 40, 50])를 생성한다.
	// 이 둘을 `append`하여 인덱스 1의 요소를 건너뛰는 효과를 낸다.
	sliceToDelete = append(sliceToDelete[:1], sliceToDelete[2:]...)
	fmt.Println("인덱스 1 (20) 삭제 후:", sliceToDelete) // 결과: [10 30 40 50]

	// 중간 요소 삭제 (인덱스 2의 값 40 삭제)
	sliceToDelete = append(sliceToDelete[:2], sliceToDelete[3:]...)
	fmt.Println("인덱스 2 (40) 삭제 후:", sliceToDelete) // 결과: [10 30 50]

	// 마지막 요소 삭제
	// 슬라이스의 길이를 1 줄여서 마지막 요소를 제외한다.
	sliceToDelete = sliceToDelete[:len(sliceToDelete)-1]
	fmt.Println("마지막 요소 (50) 삭제 후:", sliceToDelete) // 결과: [10 30]

	// 특정 값 삭제 (예: 30 삭제)
	// Go는 Java의 `remove(Object)`와 같은 직접적인 값 삭제 메서드를 제공하지 않는다.
	// 순회하면서 해당 값을 제외하고 새 슬라이스를 만들거나, 위와 같은 슬라이싱 기법을 사용해야 한다.
	// 여기서는 예시를 위해 간단한 루프를 사용한다.
	sliceByValue := []int{10, 20, 30, 20, 40}
	valToRemove := 20
	// 새 슬라이스를 생성하고, 삭제할 값을 제외한 요소들만 추가한다.
	newSlice := make([]int, 0, len(sliceByValue))
	for _, v := range sliceByValue {
		if v != valToRemove {
			newSlice = append(newSlice, v)
		}
	}
	sliceByValue = newSlice
	fmt.Println("값 20 삭제 후:", sliceByValue) // 결과: [10 30 40]
}
