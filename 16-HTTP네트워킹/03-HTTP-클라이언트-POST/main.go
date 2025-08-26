package main

import (
	"fmt"
	"io" // ioutil 대신 io 패키지 사용
	"log"
	"net/http"
	"net/url" // FormData 생성을 위해 import
)

func main() {
	postUrl := "https://httpbin.org/post"
	// 1. 전송할 FormData 데이터 생성
	formData := url.Values{}
	formData.Set("name", "Gemini")
	formData.Set("level", "100")

	// 2. http.PostForm 함수로 POST 요청
	// 이 함수는 자동으로 Content-Type을 application/x-www-form-urlencoded 로 설정합니다.
	resp, err := http.PostForm(postUrl, formData)
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// 3. 응답 본문 읽기 및 출력
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Println(string(body))
}
