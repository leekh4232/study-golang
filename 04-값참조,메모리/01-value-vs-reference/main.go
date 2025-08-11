package main

import "fmt"

func main() {
	// 1.1 값 타입 (int)
	fmt.Println("---", "1.1 값 타입 (int)", "---")
	num := 10
	fmt.Println("원본 num:", num) // 10
	modifyInt(num)
	fmt.Println("modifyInt 호출 후 num:", num) // 10 (변화 없음)

	// 1.2 값 타입 (배열)
	fmt.Println("\n---", "1.2 값 타입 (배열)", "---")
	arr := [3]int{1, 2, 3}
	fmt.Println("원본 arr:", arr) // [1 2 3]
	modifyArray(arr)
	fmt.Println("modifyArray 호출 후 arr:", arr) // [1 2 3] (변화 없음)

	// 1.3 참조 타입처럼 동작하는 타입 (슬라이스)
	fmt.Println("\n---", "1.3 참조 타입처럼 동작하는 타입 (슬라이스)", "---")
	s := []int{10, 20, 30}
	fmt.Println("원본 s:", s) // [10 20 30]
	modifySlice(s)
	fmt.Println("modifySlice 호출 후 s:", s) // [99 20 30] (변화 있음)

	// 1.4 참조 타입처럼 동작하는 타입 (맵)
	fmt.Println("\n---", "1.4 참조 타입처럼 동작하는 타입 (맵)", "---")
	m := map[string]int{"A": 1, "B": 2}
	fmt.Println("원본 m:", m) // map[A:1 B:2]
	modifyMap(m)
	fmt.Println("modifyMap 호출 후 m:", m) // map[A:99 B:2 C:3] (변화 있음)
}

// int는 값 타입이므로 복사본이 전달된다.
func modifyInt(val int) {
	val = 99
	fmt.Println("  modifyInt 내부:", val) // 99
}

// 배열은 값 타입이므로 복사본이 전달된다.
func modifyArray(arr [3]int) {
	arr[0] = 99
	fmt.Println("  modifyArray 내부:", arr) // [99 2 3]
}

// 슬라이스는 내부 배열에 대한 포인터와 길이/용량을 포함하는 구조체이다.
// 이 구조체 자체가 값으로 복사되지만, 내부 배열은 원본과 공유된다.
func modifySlice(s []int) {
	s[0] = 99                           // 내부 배열의 요소 변경
	fmt.Println("  modifySlice 내부:", s) // [99 20 30]
}

// 맵은 내부 해시 테이블에 대한 포인터를 포함하는 구조체이다.
// 이 구조체 자체가 값으로 복사되지만, 내부 해시 테이블은 원본과 공유된다.
func modifyMap(m map[string]int) {
	m["A"] = 99                       // 내부 해시 테이블의 요소 변경
	m["C"] = 3                        // 새 요소 추가
	fmt.Println("  modifyMap 내부:", m) // map[A:99 B:2 C:3]
}
