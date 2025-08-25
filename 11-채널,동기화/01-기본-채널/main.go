package main

import (
	"fmt"
	"time"
)

// 문자열을 받는 채널을 인자로 받는 함수
func worker(ch chan string) {
	fmt.Println("Worker: 작업 시작...")
	time.Sleep(500 * time.Millisecond) // 작업을 시뮬레이션하기 위해 0.5초 대기
	fmt.Println("Worker: 작업 완료")

	// 채널에 작업 완료 메시지를 보냄
	ch <- "작업이 성공적으로 끝났습니다."
}

func main() {
	// string 타입의 데이터를 주고받을 수 있는 채널 생성
	ch := make(chan string)

	// worker 함수를 고루틴으로 실행하고, 생성한 채널을 넘겨줌
	go worker(ch)

	// 채널로부터 데이터가 수신될 때까지 이 라인에서 대기(blocking)함
	msg := <-ch

	// 데이터가 수신되면 대기가 풀리고 다음 코드가 실행됨
	fmt.Printf("Main: 수신된 메시지 - '%s'\n", msg)
}
