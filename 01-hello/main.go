// 1. 패키지 선언: 이 파일이 실행 가능한 프로그램임을 나타낸다. Java의 `public class Main { ... }`과 유사하게 프로그램의 진입점을 정의한다.
package main

// 2. 외부 패키지 임포트: `fmt` 패키지는 문자열 포맷팅, 입출력 등을 다루는 Go의 표준 라이브러리이다. Java의 `import java.util.*;`와 유사하게 필요한 기능을 가져온다.
import "fmt"

// 3. 프로그램 시작점: `main` 함수는 Go 프로그램이 시작될 때 가장 먼저 실행되는 함수이다. Java의 `public static void main(String[] args)`와 동일한 역할을 한다.
func main() {
    // 4. 콘솔에 문자열 출력: `fmt` 패키지의 `Println` 함수를 호출하여 괄호 안의 내용을 콘솔에 출력한다. Java의 `System.out.println()`과 유사하다.
    fmt.Println("Hello, Go!")
}