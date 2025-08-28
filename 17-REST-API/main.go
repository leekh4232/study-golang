package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv" // <-- 추가
	"strings" // <-- 추가
)

/** 5. PK역할을 수행할 ID변수 */
var nextID = 3

/** 1. Todo 자원을 나타내는 구조체 정의 */
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/** 2. 데이터베이스를 대신할 인메모리 슬라이스 */
var todos = []Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Build REST API", Completed: false},
}

/** 3. /todos 경로의 요청을 처리하는 핸들러 */
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	/** 4. GET 요청 처리 */
	case http.MethodGet:
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

	/** 6. POST 요청 처리 */
	case http.MethodPost:
		// 폼 데이터 파싱
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// 폼 값 추출
		title := r.PostFormValue("title")
		if title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}
		// "true"가 아니면 모두 false로 처리
		completed := r.PostFormValue("completed") == "true"

		// 새 Todo 객체 생성
		newTodo := Todo{
			ID:        nextID,
			Title:     title,
			Completed: completed,
		}
		nextID++

		// 새 Todo를 슬라이스에 추가
		todos = append(todos, newTodo)

		// 성공 응답 전송
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newTodo)

	/** 7. PUT 요청 처리 */
	case http.MethodPut:
		// URL 경로에서 ID 추출
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// 요청 파라미터인 title, completed 추출
		title := r.PostFormValue("title")
		completed := r.PostFormValue("completed") == "true"

		for i := range todos {
			if todos[i].ID == id {
				todos[i].Title = title
				todos[i].Completed = completed

				// 성공 응답 전송
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	/** 8. DELETE 요청 처리 (삭제) */
	case http.MethodDelete:
		// URL 경로에서 ID 추출
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)

				// 성공 응답 전송
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// `/todos` 라고만 적용할 경우 todos 이후의 path 파라미터를 받지 못하고 404 에러가 발생함
	http.HandleFunc("/todos/", todosHandler)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
