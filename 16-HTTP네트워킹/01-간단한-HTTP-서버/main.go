package main

import (
	"fmt"
	"log"
	"net/http"
)

// 1. 요청을 처리할 핸들러 함수 정의
func handler(w http.ResponseWriter, r *http.Request) {
	// 2. ResponseWriter에 응답 내용을 씀
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// 3. "/" 경로와 핸들러 함수를 매핑
	http.HandleFunc("/", handler)

	fmt.Println("Server starting on port 8080...")
	// 4. 8080 포트에서 서버를 시작
	log.Fatal(http.ListenAndServe(":8080", nil))
}
