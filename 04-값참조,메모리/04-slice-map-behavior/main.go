package main

import "fmt"

// 슬라이스 헤더는 값으로 전달되지만, 내부 배열은 공유된다.
func modifyAndReassignSlice(s []int) {
	fmt.Println("  modifyAndReassignSlice 내부 - 초기 s:", s, "길이:", len(s), "용량:", cap(s))
	s[0] = 999 // 내부 배열의 요소 변경 (원본에 영향)
	fmt.Println("  modifyAndReassignSlice 내부 - 요소 변경 후 s:", s)

	// s 변수 자체를 재할당 (새로운 슬라이스 헤더 생성)
	// 이 변경은 함수 내부의 s 복사본에만 적용되고, 호출자의 원본 s에는 영향을 주지 않는다.
	s = append(s, 100, 200)
	fmt.Println("  modifyAndReassignSlice 내부 - append 후 s:", s, "길이:", len(s), "용량:", cap(s))
}

// 맵 헤더는 값으로 전달되지만, 내부 해시 테이블은 공유된다.
func modifyMapInFunction(m map[string]int) {
	fmt.Println("  modifyMapInFunction 내부 - 초기 m:", m)
	m["newKey"] = 123 // 새 키-값 추가 (원본에 영향)
	delete(m, "B")    // 키 삭제 (원본에 영향)
	fmt.Println("  modifyMapInFunction 내부 - 변경 후 m:", m)
}

func main() {
	// 4.1 슬라이스 동작 방식
	fmt.Println("---", "4.1 슬라이스 동작 방식", "---")
	mySlice := []int{10, 20, 30}
	fmt.Println("main 함수 시작 mySlice:", mySlice, "길이:", len(mySlice), "용량:", cap(mySlice))
	modifyAndReassignSlice(mySlice)
	fmt.Println("main 함수 종료 mySlice:", mySlice, "길이:", len(mySlice), "용량:", cap(mySlice)) // 요소 변경은 반영, append는 반영 안됨

	// 4.2 맵 동작 방식
	fmt.Println("\n---", "4.2 맵 동작 방식", "---")
	myMap := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println("main 함수 시작 myMap:", myMap)
	modifyMapInFunction(myMap)
	fmt.Println("main 함수 종료 myMap:", myMap) // 변경 사항 모두 반영
}
