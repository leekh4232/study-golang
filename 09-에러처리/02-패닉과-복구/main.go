package main

import "fmt"

// panic으로부터 프로그램을 복구하는 함수
func handlePanic() {
	// recover() 함수를 호출
	// 만약 패닉 상태라면, recover는 panic에 전달된 값을 반환함
	a := recover()
	if a != nil {
		fmt.Println("RECOVERED:", a) // 복구 메시지 출력
	}
}

// 패닉을 발생시키는 함수
func causePanic() {
	// 함수가 종료되기 직전에 항상 handlePanic이 실행되도록 defer 설정
	defer handlePanic()

	// "심각한 에러" 메시지와 함께 패닉 발생
	panic("심각한 에러")
}

func main() {
	fmt.Println("main 함수 시작")
	causePanic() // 패닉을 발생시키는 함수 호출
	// causePanic 내부에서 panic이 발생했지만, handlePanic에 의해 복구되었으므로
	// 프로그램이 종료되지 않고 이 라인이 실행됨.
	fmt.Println("main 함수 종료")
}
