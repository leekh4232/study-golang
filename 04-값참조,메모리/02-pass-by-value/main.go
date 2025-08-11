package main

import "fmt"

// int는 값 타입이므로, 함수 내에서 변경해도 원본에 영향 없음
func increment(x int) {
	x++
	fmt.Println("  increment 함수 내부 x:", x)
}

// 배열은 값 타입이므로, 함수 내에서 변경해도 원본에 영향 없음
func changeArray(arr [3]int) {
	arr[0] = 100
	fmt.Println("  changeArray 함수 내부 arr:", arr)
}

// 구조체는 값 타입이므로, 함수 내에서 변경해도 원본에 영향 없음
type Player struct {
	Name  string
	Level int
}

func levelUp(p Player) {
	p.Level++
	fmt.Println("  levelUp 함수 내부 Player:", p)
}

// 슬라이스는 헤더가 값으로 전달되지만, 내부 배열은 공유된다.
// 따라서 슬라이스 요소를 변경하면 원본에 영향 있음.
func appendAndModifySlice(s []int) {
	s[0] = 999            // 원본 슬라이스 요소 변경
	s = append(s, 40, 50) // s 변수 자체를 재할당 (원본에는 영향 없음)
	fmt.Println("  appendAndModifySlice 함수 내부 s:", s)
}

func main() {
	// 2.1 기본 타입의 값 전달
	fmt.Println("---", "2.1 기본 타입의 값 전달", "---")
	a := 5
	fmt.Println("main 함수 시작 a:", a) // 5
	increment(a)
	fmt.Println("main 함수 종료 a:", a) // 5 (변화 없음)

	// 2.2 배열의 값 전달
	fmt.Println("\n---", "2.2 배열의 값 전달", "---")
	myArr := [3]int{1, 2, 3}
	fmt.Println("main 함수 시작 myArr:", myArr) // [1 2 3]
	changeArray(myArr)
	fmt.Println("main 함수 종료 myArr:", myArr) // [1 2 3] (변화 없음)

	// 2.3 구조체의 값 전달
	fmt.Println("\n---", "2.3 구조체의 값 전달", "---")
	p := Player{Name: "Hero", Level: 1}
	fmt.Println("main 함수 시작 Player:", p) // {Hero 1}
	levelUp(p)
	fmt.Println("main 함수 종료 Player:", p) // {Hero 1} (변화 없음)

	// 2.4 슬라이스의 값 전달 (내부 데이터 공유)
	fmt.Println("\n---", "2.4 슬라이스의 값 전달 (내부 데이터 공유)", "---")
	mySlice := []int{10, 20, 30}
	fmt.Println("main 함수 시작 mySlice:", mySlice) // [10 20 30]
	appendAndModifySlice(mySlice)
	fmt.Println("main 함수 종료 mySlice:", mySlice) // [999 20 30] (요소 변경은 반영, append는 반영 안됨)
}
