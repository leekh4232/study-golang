package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// 1. http.Get 함수로 GET 요청 전송
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	// 2. 함수 종료 시점에 응답 본문을 닫도록 예약
	defer resp.Body.Close()

	// 3. 응답 본문을 모두 읽음
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 4. 읽은 본문을 문자열로 변환하여 출력
	fmt.Println(string(body))
}
