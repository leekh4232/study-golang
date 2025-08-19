package main

import "fmt"

func main() {
	/** 배열 선언 및 초기화 */

	// Go 배열: 고정된 크기. Java의 `int[] arr = new int[3];`와 유사하지만, Go 배열은 값 타입이다.
	// 배열의 크기는 타입의 일부이므로, [3]int와 [4]int는 다른 타입이다.
	var arr1 = [3]int{1, 2, 3}
	fmt.Println("배열 arr1:", arr1)

	// Go 배열 리터럴: `...`을 사용하여 컴파일러가 초기화 값의 개수에 따라 배열 크기를 자동 추론하게 한다.
	// Java의 `{}` 초기화와 유사.
	arr2 := [...]string{"apple", "banana", "cherry"}
	fmt.Println("배열 arr2:", arr2)

	// 배열은 값 타입(Value Type): 함수에 전달 시 복사됨. Java의 기본 타입(int, boolean 등)과 유사.
	// 따라서 `modifyArray` 함수 내에서 `arr_copy`를 변경해도 원본 `arr1`에는 영향을 주지 않는다.
	modifyArray(arr1)
	fmt.Println("modifyArray 호출 후 arr1 (변화 없음):", arr1)

	/** 슬라이스 선언 및 초기화 */

	// Go 슬라이스 리터럴: 동적 크기. Java의 `ArrayList<Integer> list = new ArrayList<>(Arrays.asList(4, 5, 6));`와 유사.
	// 슬라이스는 내부적으로 배열을 참조하며, 길이(len)와 용량(cap)을 가진다.
	slice1 := []int{4, 5, 6}
	fmt.Println("슬라이스 slice1:", slice1)

	// `make` 함수를 이용한 슬라이스 생성: Java의 `new ArrayList<>(initialCapacity);`와 유사.
	// `make([]T, length, capacity)`
	// `length`: 슬라이스의 초기 길이 (요소는 제로 값으로 초기화된다).
	// `capacity`: 슬라이스의 내부 배열이 가질 수 있는 최대 길이 (생략 시 `length`와 동일).
	slice2 := make([]int, 5) // 길이 5, 용량 5인 슬라이스 (모든 요소는 제로 값으로 초기화)
	fmt.Println("슬라이스 slice2 (make로 생성):", slice2)

	slice3 := make([]string, 0, 10) // 길이 0, 용량 10인 슬라이스. 요소를 추가할 때 용량 내에서 효율적으로 확장된다.
	fmt.Println("슬라이스 slice3 (길이 0, 용량 10):", slice3)

	// 슬라이스는 참조 타입(Reference Type)처럼 동작: 함수에 전달 시 내부 배열을 공유.
	// Java에서 객체(예: ArrayList)를 함수에 전달하여 객체 내부를 변경하는 것과 유사.
	modifySlice(slice1)
	fmt.Println("modifySlice 호출 후 slice1 (변화 있음):", slice1)

	// nil 슬라이스: 선언만 하고 초기화하지 않은 슬라이스 (길이 0, 용량 0, 내부 배열 없음).
	// Java의 `null` 참조와 유사하지만, Go의 nil 슬라이스는 `len`과 `cap` 함수를 사용할 수 있으며, `append`도 가능하다.
	var nilSlice []int
	fmt.Printf("nilSlice: %v, 길이: %d, 용량: %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil { // nil 슬라이스는 `nil`과 비교할 수 있다.
		fmt.Println("nilSlice는 nil이다.")
	}
}

// 배열을 수정하는 함수 (배열은 값 타입이므로 원본은 변경되지 않음)
// 함수 인자로 배열이 전달될 때, 배열 전체가 복사된다.
func modifyArray(arr [3]int) {
	arr[0] = 99 // 복사된 배열의 첫 번째 요소를 변경
	fmt.Println("modifyArray 내부:", arr)
}

// 슬라이스를 수정하는 함수 (슬라이스는 내부 배열을 공유하므로 원본이 변경될 수 있음)
// 함수 인자로 슬라이스가 전달될 때, 슬라이스 헤더(포인터, 길이, 용량)가 복사된다.
// 이 복사된 헤더의 포인터는 원본 슬라이스의 내부 배열을 가리킨다.
func modifySlice(s []int) {
	s[0] = 99 // 슬라이스 헤더의 포인터를 통해 내부 배열의 첫 번째 요소를 변경. 이는 원본 슬라이스에도 영향을 준다.
	fmt.Println("modifySlice 내부:", s)
}
