package main

import (
	"errors"
	"fmt"
)

// 두 수를 나누는 함수. 실패 시 error를 반환함.
func divide(a, b float64) (float64, error) {
	// 나누는 수가 0인 경우
	if b == 0 {
		// errors.New 함수로 새로운 에러 메시지를 생성하여 반환
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	// 성공한 경우, 결과와 함께 nil을 반환 (에러가 없다는 의미)
	return a / b, nil
}

func main() {
	// 성공 사례
	result, err := divide(10, 2)
	if err != nil { // 에러가 발생했는지 확인
		fmt.Printf("에러 발생: %s\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// 실패 사례
	result, err = divide(10, 0)
	if err != nil { // 에러가 발생했는지 확인
		fmt.Printf("에러 발생: %s\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}
}
