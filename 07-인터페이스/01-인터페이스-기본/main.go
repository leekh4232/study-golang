package main

import "fmt"

// Notifier 인터페이스를 정의.
// 이 인터페이스는 string 타입의 메시지를 받는 Send 메서드 하나를 가지고 있음.
// 이 메서드를 구현하는 모든 타입은 Notifier로 간주될 수 있음.
type Notifier interface {
	Send(message string)
}

// EmailNotifier 구조체를 정의. 이메일 수신자 주소를 필드로 가짐.
type EmailNotifier struct {
	Recipient string
}

// EmailNotifier 타입에 대한 Send 메서드를 구현.
// 이 메서드는 Notifier 인터페이스의 요구사항을 만족시킴.
// (e *EmailNotifier)는 포인터 리시버로, 메서드 내에서 구조체 필드 변경이 가능함을 의미.
func (e *EmailNotifier) Send(message string) {
	// fmt.Printf를 사용하여 포맷팅된 문자열을 콘솔에 출력.
	fmt.Printf("'%s' 주소로 이메일 발송: %s", e.Recipient, message)
}

func main() {
	// Notifier 타입의 변수 n을 선언.
	// 인터페이스는 그 자체로 인스턴스화될 수 없으며, 인터페이스를 구현한 구체적인 타입의 값을 담는 컨테이너 역할을 함.
	var n Notifier

	// EmailNotifier 구조체의 인스턴스를 생성하고 그 메모리 주소를 n에 할당.
	// EmailNotifier가 Send 메서드를 구현했기 때문에 Notifier 타입 변수에 할당이 가능함.
	// 이것이 Go의 '암시적 인터페이스 구현'임.
	n = &EmailNotifier{Recipient: "test@example.com"}

	// 인터페이스 변수 n을 통해 Send 메서드를 호출.
	// Go 런타임은 n이 실제로 가리키는 EmailNotifier의 Send 메서드를 찾아 실행해줌.
	n.Send("안녕하세요! Go 인터페이스 테스트입니다.")
}
