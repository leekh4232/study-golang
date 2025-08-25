package main

import (
	"fmt"
	"io" // 입출력 프리미티브를 제공하는 패키지
	"os" // 운영체제 인터페이스 패키지
)

func main() {
	// 1. 파일 열기: os.Open()은 읽기 전용으로 파일을 엶
	//    Java의 new FileInputStream("example.txt")와 유사
	file, err := os.Open("../example.txt")
	if err != nil {
		fmt.Println("파일 열기 실패:", err)
		return // Java의 FileNotFoundException 처리와 유사
	}
	// 2. defer로 파일 닫기 보장
	defer file.Close()

	// 3. 파일 전체 내용 읽기: io.ReadAll()은 파일의 모든 내용을 메모리로 읽음
	//    작은 파일에 적합하며, Java의 Files.readAllBytes()와 유사
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("파일 읽기 실패:", err)
		return
	}

	// 4. 읽은 내용 출력: 바이트 슬라이스를 문자열로 변환
	//    Java의 new String(bytes, StandardCharsets.UTF_8)와 유사
	fmt.Println("파일 내용:")
	fmt.Printf("--- 시작 ---\n%s\n--- 끝 ---\n", string(content))

	// 5. 파일 크기 정보 출력
	fmt.Printf("총 %d 바이트를 읽었습니다.\n", len(content))
}
