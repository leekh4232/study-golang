package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 두 숫자와 연산자를 받아 계산 결과를 반환하는 함수
// Go 함수는 다중 반환 값을 지원한다. (Java는 객체나 배열로 묶어서 반환해야 함)
func calculate(num1, num2 float64, operator string) (float64, error) {
	var result float64
	var err error

	// Go의 switch 문은 Java보다 유연하다. break를 명시하지 않아도 자동으로 다음 case로 넘어가지 않는다.
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("0으로 나눌 수 없습니다.") // Go의 에러 반환 방식
		}
		result = num1 / num2
	default:
		err = fmt.Errorf("유효하지 않은 연산자입니다: %s", operator)
	}
	return result, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 첫 번째 숫자 입력
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	input1, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
		return
	}

	// 연산자 입력
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	// 두 번째 숫자 입력
	fmt.Print("두 번째 숫자를 입력하세요: ")
	input2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
		return
	}

	// 계산 함수 호출 및 결과 처리
	// Go의 다중 반환 값 처리 (Java는 try-catch로 예외 처리)
	result, calcErr := calculate(num1, num2, operator)
	if calcErr != nil {
		fmt.Println("에러:", calcErr)
	} else {
		fmt.Printf("결과: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}

	// 종료 대기 메시지 및 입력 대기
	fmt.Print("종료하려면 아무 키나 누르세요...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
