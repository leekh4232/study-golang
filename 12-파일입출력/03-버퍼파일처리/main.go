package main

import (
	"bufio" // 버퍼링된 I/O 연산을 제공하는 패키지
	"fmt"
	"os"
)

func main() {
	// 1. 먼저 테스트용 대용량 파일 생성
	createLargeFile()

	// 2. 버퍼를 사용한 라인별 읽기
	readFileLineByLine()
}

// 테스트용 파일 생성 함수
func createLargeFile() {
	file, err := os.Create("../large_file.txt")
	if err != nil {
		fmt.Println("대용량 파일 생성 실패:", err)
		return
	}
	defer file.Close()

	// bufio.Writer를 사용하여 버퍼링된 쓰기 (Java의 BufferedWriter와 유사)
	writer := bufio.NewWriter(file)

	// 1000줄의 텍스트 데이터 생성
	for i := 1; i <= 1000; i++ {
		line := fmt.Sprintf("이것은 %d번째 줄입니다. 버퍼링된 파일 처리를 테스트합니다.\n", i)
		// WriteString은 문자열을 바로 쓸 수 있는 편의 메서드
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println("라인 쓰기 실패:", err)
			return
		}
	}

	// 버퍼의 내용을 파일에 플러시 (Java의 BufferedWriter.flush()와 동일)
	err = writer.Flush()
	if err != nil {
		fmt.Println("버퍼 플러시 실패:", err)
		return
	}

	fmt.Println("테스트용 대용량 파일 생성 완료")
}

// 라인별 파일 읽기 함수
func readFileLineByLine() {
	file, err := os.Open("../large_file.txt")
	if err != nil {
		fmt.Println("파일 열기 실패:", err)
		return
	}
	defer file.Close()

	// bufio.Scanner를 사용한 라인별 읽기 (Java의 Scanner와 유사)
	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Scan() 메서드는 다음 토큰(기본적으로 라인)이 있으면 true 반환
	// Java의 Scanner.hasNextLine()과 유사한 패턴
	for scanner.Scan() {
		lineCount++
		line := scanner.Text() // 현재 라인의 텍스트 반환 (개행문자 제외)

		// 처음 5줄과 마지막 5줄만 출력 (메모리 효율성 시연)
		if lineCount <= 5 || lineCount > 995 {
			fmt.Printf("라인 %d: %s\n", lineCount, line)
		} else if lineCount == 6 {
			fmt.Println("... (중간 라인들 생략) ...")
		}
	}

	// 스캔 과정에서 발생한 에러 확인 (EOF는 정상적인 종료이므로 에러가 아님)
	if err := scanner.Err(); err != nil {
		fmt.Println("파일 스캔 중 에러 발생:", err)
		return
	}

	fmt.Printf("\n총 %d줄을 처리했습니다.\n", lineCount)
}
