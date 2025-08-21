package myapp // 패키지를 'myapp'으로 선언

// 두 정수의 합을 반환하는 함수
// 함수 이름이 대문자 'A'로 시작하므로 외부 패키지에서 호출할 수 있음 (Public)
func Add(a int, b int) int {
	return a + b
}

// 두 정수의 차를 반환하는 함수
// 함수 이름이 소문자 's'로 시작하므로 외부에서 호출할 수 없음 (Private)
func subtract(a int, b int) int {
	return a - b
}
