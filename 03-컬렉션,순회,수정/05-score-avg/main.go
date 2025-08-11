package main

import "fmt"

func main() {
	/** 맵을 이용한 점수 저장 */
	// Go 맵을 이용하여 학생 이름(string)을 키로, 점수(int)를 값으로 저장한다.
	// Java의 `HashMap<String, Integer> scores = new HashMap<>();`와 유사하다.
	scores := map[string]int{"Alice": 90, "Bob": 85, "Charlie": 95}

	sum := 0 // 총점을 저장할 변수 초기화

	/** `for-range`를 이용한 맵 순회 및 합계 계산 */
	// 맵을 `for-range`로 순회하여 각 학생의 점수를 `sum`에 더한다.
	// 키(`name`)는 필요 없으므로 `_`로 무시한다.
	// Java의 `for (int score : scores.values()) { sum += score; }`와 유사하다.
	for _, score := range scores {
		sum += score
	}

	/** 평균 계산 */
	// Go는 정수 나눗셈 시 소수점 이하를 버리므로, 정확한 평균을 위해 명시적 타입 변환이 필요하다.
	// `float64(sum)`은 `sum`을 `float64` 타입으로 변환한다.
	// `len(scores)`는 맵의 요소 개수(학생 수)를 반환한다.
	// Java의 `double avg = (double) sum / scores.size();`와 유사하다.
	avg := float64(sum) / float64(len(scores))
	// `fmt.Printf`를 사용하여 포맷팅된 문자열을 출력한다.
	// `%.2f`는 소수점 이하 두 자리까지 출력하도록 지정한다.
	fmt.Printf("총점: %d, 평균: %.2f\n", sum, avg)
}
