package main

import (
	"encoding/json" // JSON 인코딩/디코딩을 위한 패키지
	"fmt"
	"os"
)

// User 구조체: JSON으로 변환될 데이터 구조
// Java의 POJO 클래스와 유사하지만, 태그를 통해 JSON 필드명을 지정할 수 있음
type User struct {
	ID       int      `json:"id"`        // JSON에서 "id" 필드로 매핑
	Name     string   `json:"name"`      // JSON에서 "name" 필드로 매핑
	Email    string   `json:"email"`     // JSON에서 "email" 필드로 매핑
	Age      int      `json:"age"`       // JSON에서 "age" 필드로 매핑
	Hobbies  []string `json:"hobbies"`   // JSON 배열로 매핑
	IsActive bool     `json:"is_active"` // JSON에서 "is_active" 필드로 매핑
}

// UserConfig 구조체: 여러 사용자를 포함하는 설정 파일 형태
type UserConfig struct {
	Version     string `json:"version"`      // 설정 파일 버전
	LastUpdated string `json:"last_updated"` // 마지막 업데이트 시간
	Users       []User `json:"users"`        // 사용자 목록
}

func main() {
	// 1. 샘플 데이터 생성: Java의 생성자 패턴과 유사하지만 더 간결함
	config := UserConfig{
		Version:     "1.0",
		LastUpdated: "2024-11-01T09:00:00Z",
		Users: []User{
			{
				ID:       1,
				Name:     "김철수",
				Email:    "kim@example.com",
				Age:      25,
				Hobbies:  []string{"독서", "프로그래밍", "영화감상"},
				IsActive: true,
			},
			{
				ID:       2,
				Name:     "이영희",
				Email:    "lee@example.com",
				Age:      30,
				Hobbies:  []string{"요리", "여행"},
				IsActive: false,
			},
			{
				ID:       3,
				Name:     "박민수",
				Email:    "park@example.com",
				Age:      28,
				Hobbies:  []string{"운동", "게임", "음악"},
				IsActive: true,
			},
		},
	}

	// 2. 구조체를 JSON으로 변환: json.MarshalIndent()는 가독성 좋은 JSON 생성
	//    Java의 ObjectMapper.writeValueAsString()과 유사
	//    첫 번째 매개변수: 변환할 객체
	//    두 번째 매개변수: 각 줄의 접두사 (보통 빈 문자열)
	//    세 번째 매개변수: 들여쓰기 문자 (보통 공백 2개 또는 탭)
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("JSON 변환 실패:", err)
		return
	}

	// 3. JSON 데이터를 파일에 쓰기: os.WriteFile()은 한 번에 파일 쓰기 완료
	//    Java의 Files.write()와 유사하지만 더 간단함
	err = os.WriteFile("../users.json", jsonData, 0644) // 0644는 파일 권한 (rw-r--r--)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
		return
	}

	fmt.Println("JSON 파일 생성 완료: users.json")
	fmt.Printf("파일 크기: %d 바이트\n", len(jsonData))

	// 4. 생성된 JSON 내용의 일부를 콘솔에 출력
	fmt.Println("\n생성된 JSON 내용:")
	fmt.Println(string(jsonData))
}
