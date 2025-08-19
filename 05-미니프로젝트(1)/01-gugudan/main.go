package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("구구단을 출력할 숫자를 입력하세요 (2-9): ")
	reader := bufio.NewReader(os.Stdin) // 표준 입력 리더 생성
	input, _ := reader.ReadString('\n') // 한 줄 읽기
	input = strings.TrimSpace(input)    // 공백 제거

	// 문자열을 정수로 변환
	// Java의 `Integer.parseInt()`와 유사하지만, Go는 에러도 함께 반환한다.
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
		return
	}

	if num < 2 || num > 9 {
		fmt.Println("2에서 9 사이의 숫자를 입력해주세요.")
		return
	}

	fmt.Printf("--- %d단 ---\n", num)
	// Go의 for 반복문 (Java의 for 루프와 유사)
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}

	// 종료 대기 메시지 및 입력 대기
	fmt.Print("종료하려면 아무 키나 누르세요...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
