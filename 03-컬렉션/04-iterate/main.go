package main

import "fmt"

func main() {
	/** 슬라이스 순회 */

	// for-range는 인덱스와 값을 반환한다.
	// 값은 복사본이므로, 순회 중 값을 변경해도 원본 슬라이스에는 영향을 주지 않는다.
	// Java의 향상된 for문(for-each)과 유사하지만, Go는 인덱스도 함께 얻을 수 있다.
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("--- 슬라이스 순회 ---")
	for i, num := range numbers {
		fmt.Printf("인덱스: %d, 값: %d\n", i, num)
		// num = 99 // num은 복사본이므로 원본 numbers에는 영향을 주지 않는다.
	}
	fmt.Println("순회 후 슬라이스 (변화 없음):", numbers)

	// 값만 필요한 경우 인덱스를 _로 무시할 수 있다.
	// Java의 `for (int num : numbers)`와 동일한 효과.
	fmt.Println("--- 슬라이스 값만 순회 ---")
	for _, num := range numbers {
		fmt.Println("값:", num)
	}

	/** 맵 순회 */

	// 맵은 순서가 보장되지 않는다. 순회할 때마다 순서가 다를 수 있다.
	// for-range는 키와 값을 반환한다.
	// Java의 `for (Map.Entry<String, Integer> entry : scores.entrySet())`와 유사.
	scores := map[string]int{"Alice": 90, "Bob": 85, "Charlie": 95}
	fmt.Println("--- 맵 순회 ---")
	for name, score := range scores {
		fmt.Printf("학생: %s, 점수: %d\n", name, score)
	}

	// 키만 필요한 경우 값을 _로 무시할 수 있다.
	// Java의 `for (String name : scores.keySet())`와 유사.
	fmt.Println("--- 맵 키만 순회 ---")
	for name := range scores {
		fmt.Println("학생:", name)
	}

	/** 문자열 순회 */

	// 문자열을 for-range로 순회하면 유니코드 코드 포인트(rune)를 반환한다.
	// 인덱스는 바이트 오프셋이다. Java는 `String.charAt(index)`로 char를 얻지만, 유니코드 처리는 별도.
	koreanString := "안녕하세요"
	fmt.Println("--- 문자열 순회 ---")
	for i, r := range koreanString {
		fmt.Printf("인덱스: %d, 룬(문자): %c, 유니코드 값: %U\n", i, r, r)
	}

	/** 배열 순회 (슬라이스와 유사) */

	arr := [3]string{"A", "B", "C"}
	fmt.Println("--- 배열 순회 ---")
	for i, val := range arr {
		fmt.Printf("인덱스: %d, 값: %s\n", i, val)
	}
}
