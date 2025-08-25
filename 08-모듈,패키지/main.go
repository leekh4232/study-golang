package main

import (
	"fmt"
	"hello-go/myapp" // hello-go 모듈에 속한 myapp 패키지를 import
)

func main() {
	// myapp 패키지의 Add 함수 호출
	sum := myapp.Add(5, 3)
	fmt.Printf("5 + 3 = %d", sum)

	// 아래 코드는 컴파일 에러를 발생시킴
	// myapp 패키지의 subtract 함수는 소문자로 시작하여 외부에서 접근할 수 없기 때문
	// diff := myapp.subtract(5, 3)
	// fmt.Printf("5 - 3 = %d", diff)
}
