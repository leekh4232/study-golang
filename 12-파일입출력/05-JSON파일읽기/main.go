package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 이전 예제와 동일한 구조체 정의
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Age      int      `json:"age"`
	Hobbies  []string `json:"hobbies"`
	IsActive bool     `json:"is_active"`
}

type UserConfig struct {
	Version     string `json:"version"`
	LastUpdated string `json:"last_updated"`
	Users       []User `json:"users"`
}

func main() {
	// 1. JSON 파일 읽기: os.ReadFile()은 파일 전체를 바이트 슬라이스로 읽음
	//    Java의 Files.readAllBytes()와 동일한 기능
	jsonData, err := os.ReadFile("../users.json")
	if err != nil {
		fmt.Println("파일 읽기 실패:", err)
		fmt.Println("먼저 JSON 파일 쓰기 예제를 실행하세요.")
		return
	}

	// 2. JSON 데이터를 구조체로 변환: json.Unmarshal()로 역직렬화
	//    Java의 ObjectMapper.readValue()와 유사
	//    두 번째 매개변수는 포인터여야 함 (값을 직접 수정하기 위해)
	var config UserConfig
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		fmt.Println("JSON 파싱 실패:", err)
		return
	}

	// 3. 읽어온 데이터 출력 및 분석
	fmt.Println("=== JSON 파일 읽기 성공 ===")
	fmt.Printf("설정 파일 버전: %s\n", config.Version)
	fmt.Printf("마지막 업데이트: %s\n", config.LastUpdated)
	fmt.Printf("총 사용자 수: %d명\n\n", len(config.Users))

	// 4. 각 사용자 정보 출력
	fmt.Println("=== 사용자 목록 ===")
	for i, user := range config.Users {
		fmt.Printf("[%d] 사용자 정보:\n", i+1)
		fmt.Printf("    ID: %d\n", user.ID)
		fmt.Printf("    이름: %s\n", user.Name)
		fmt.Printf("    이메일: %s\n", user.Email)
		fmt.Printf("    나이: %d세\n", user.Age)
		fmt.Printf("    활성 상태: %t\n", user.IsActive)
		fmt.Printf("    취미: %v\n", user.Hobbies) // %v는 값의 기본 형식으로 출력
		fmt.Println()
	}
}
