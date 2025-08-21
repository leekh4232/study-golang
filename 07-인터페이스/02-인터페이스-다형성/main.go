package main

import "fmt"

// Notifier 인터페이스는 변경 없이 그대로 사용.
type Notifier interface {
	Send(message string)
}

// EmailNotifier 구조체와 Send 메서드도 그대로 사용.
type EmailNotifier struct {
	Recipient string
}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("'%s' 주소로 이메일 발송: %s\n", e.Recipient, message)
}

// 새로운 알림 방식인 SMSNotifier 구조체를 정의.
type SMSNotifier struct {
	PhoneNumber string
}

// SMSNotifier 타입에 대한 Send 메서드를 구현.
// 이 또한 Notifier 인터페이스의 요구사항을 만족시킴.
func (s *SMSNotifier) Send(message string) {
	fmt.Printf("'%s' 번호로 SMS 발송: %s\n", s.PhoneNumber, message)
}

// 다형성을 보여주는 핵심 함수.
// 매개변수로 Notifier 인터페이스 타입을 받음.
// 이 함수는 전달된 실제 타입이 EmailNotifier인지 SMSNotifier인지 신경쓰지 않고,
// 오직 Send 메서드를 호출할 수 있다는 사실에만 의존함.
func SendNotification(n Notifier, message string) {
	fmt.Println("알림을 전송합니다...")
	n.Send(message)
}

func main() {
	// EmailNotifier 인스턴스를 생성.
	emailNotifier := &EmailNotifier{Recipient: "dev@hossam.com"}
	// SMSNotifier 인스턴스를 생성.
	smsNotifier := &SMSNotifier{PhoneNumber: "010-1234-5678"}

	// 동일한 SendNotification 함수에 서로 다른 타입의 값을 전달.
	// 인터페이스 덕분에 함수는 유연하게 두 타입을 모두 처리할 수 있음.
	SendNotification(emailNotifier, "새로운 기능이 배포되었습니다.")
	fmt.Println("---------------")
	SendNotification(smsNotifier, "서버 점검이 30분 후에 시작됩니다.")
}
