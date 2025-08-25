package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// AppConfig 구조체: 애플리케이션 설정을 표현
type AppConfig struct {
	AppName     string          `json:"app_name"`
	Version     string          `json:"version"`
	Debug       bool            `json:"debug"`
	Port        int             `json:"port"`
	Database    DatabaseConfig  `json:"database"`
	Features    map[string]bool `json:"features"` // 기능 토글
	LastUpdated time.Time       `json:"last_updated"`
}

// DatabaseConfig 구조체: 데이터베이스 설정 (중첩 구조체 예시)
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // omitempty: 빈 값이면 JSON에서 제외
}

const configFileName = "../app_config.json"

func main() {
	// 1. 기본 설정 파일이 없으면 생성
	if !fileExists(configFileName) {
		fmt.Println("설정 파일이 없습니다. 기본 설정 파일을 생성합니다...")
		createDefaultConfig()
	}

	// 2. 설정 파일 읽기
	config, err := loadConfig()
	if err != nil {
		fmt.Println("설정 파일 로드 실패:", err)
		return
	}

	// 3. 현재 설정 출력
	printConfig(config)

	// 4. 사용자 입력을 통한 설정 수정 시뮬레이션
	fmt.Println("\n=== 설정 수정 ===")
	config.Debug = !config.Debug     // 디버그 모드 토글
	config.Features["new_ui"] = true // 새로운 기능 활성화
	config.LastUpdated = time.Now()  // 수정 시간 업데이트

	// 5. 수정된 설정 저장
	err = saveConfig(config)
	if err != nil {
		fmt.Println("설정 파일 저장 실패:", err)
		return
	}

	fmt.Println("설정이 성공적으로 업데이트되었습니다.")
}

// 파일 존재 여부 확인 함수
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err) // 파일이 존재하지 않는 에러가 아니면 파일이 존재함
}

// 기본 설정 파일 생성 함수
func createDefaultConfig() error {
	// 기본 설정값들 정의
	defaultConfig := AppConfig{
		AppName: "Go File I/O 실습 앱",
		Version: "1.0.0",
		Debug:   false,
		Port:    8080,
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Name:     "myapp_db",
			Username: "admin",
			Password: "", // 보안상 빈 값으로 설정
		},
		Features: map[string]bool{
			"logging":    true,
			"monitoring": false,
			"caching":    true,
			"new_ui":     false,
		},
		LastUpdated: time.Now(),
	}

	return saveConfig(defaultConfig)
}

// 설정 파일 로드 함수
func loadConfig() (AppConfig, error) {
	var config AppConfig

	// 파일 읽기
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return config, fmt.Errorf("파일 읽기 실패: %w", err) // 에러 래핑
	}

	// JSON 파싱
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("JSON 파싱 실패: %w", err)
	}

	return config, nil
}

// 설정 파일 저장 함수
func saveConfig(config AppConfig) error {
	// JSON으로 변환 (가독성을 위해 들여쓰기 포함)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 변환 실패: %w", err)
	}

	// 파일에 쓰기 (0644 권한: 소유자는 읽기/쓰기, 그룹과 기타는 읽기만)
	err = os.WriteFile(configFileName, data, 0644)
	if err != nil {
		return fmt.Errorf("파일 쓰기 실패: %w", err)
	}

	return nil
}

// 설정 정보 출력 함수
func printConfig(config AppConfig) {
	fmt.Println("=== 현재 설정 ===")
	fmt.Printf("앱 이름: %s\n", config.AppName)
	fmt.Printf("버전: %s\n", config.Version)
	fmt.Printf("디버그 모드: %t\n", config.Debug)
	fmt.Printf("포트: %d\n", config.Port)
	fmt.Printf("마지막 업데이트: %s\n", config.LastUpdated.Format("2006-01-02 15:04:05"))

	fmt.Println("\n데이터베이스 설정:")
	fmt.Printf("  호스트: %s\n", config.Database.Host)
	fmt.Printf("  포트: %d\n", config.Database.Port)
	fmt.Printf("  데이터베이스명: %s\n", config.Database.Name)
	fmt.Printf("  사용자명: %s\n", config.Database.Username)

	fmt.Println("\n기능 설정:")
	for feature, enabled := range config.Features {
		status := "비활성화"
		if enabled {
			status = "활성화"
		}
		fmt.Printf("  %s: %s\n", feature, status)
	}
}
