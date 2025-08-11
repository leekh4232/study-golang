package main

import "fmt"

func main() {
	// 6.1 기본 for문 (Java의 `for (int i = 0; i < 3; i++)`와 동일)
	fmt.Println("--- 기본 for문 ---")
	for i := 0; i < 3; i++ { // 초기화; 조건; 증감문
		fmt.Println(i)
	}

	// 6.2 `while`문처럼 사용 (Java의 `while (count < 3)`와 동일)
	// Go에는 별도의 `while` 키워드가 없으며, `for` 조건문만 사용하여 `while` 루프를 구현한다.
	fmt.Println("--- while처럼 사용 ---")
	count := 0
	for count < 3 { // 조건문만 사용
		fmt.Println(count)
		count++
	}

	// 6.3 `for-range`문 (Java의 향상된 for문 `for (Type item : collection)`과 유사)
	// 배열, 슬라이스, 맵, 문자열, 채널 등을 순회할 때 사용한다.
	// 슬라이스/배열/문자열의 경우 `index`와 `value`를 반환한다.
	fmt.Println("--- for-range문 ---")
	items := []string{"물약", "검", "방패"} // 문자열 슬라이스 선언
	for index, item := range items {   // `index`는 요소의 인덱스, `item`은 요소의 값
		fmt.Printf("%d: %s\n", index, item)
	}

	// 6.4 무한 루프 (Java의 `while (true)` 또는 `for (;;)`)와 동일
	// `for {}` 형태로 사용한다.
	// for {
	//     fmt.Println("무한 루프!")
	//     time.Sleep(time.Second) // 1초 대기
	// }
}
