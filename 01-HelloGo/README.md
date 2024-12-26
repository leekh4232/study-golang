# Hello Go

## #01. Go 언어 소개

- Go(고 또는 Golang)는 구글에서 개발한 오픈소스 프로그래밍 언어
- 2009년에 처음 공개
- 성능, 간결함, 병렬 처리에 초점을 맞추어 설계됨.
- 특히, 대규모 네트워크 애플리케이션과 클라우드 기반 환경에서의 사용을 염두에 두고 만들어짐

### Java와의 차이점

| **항목**| **Go (Golang)** | **Java** |
|--|--|--|
| **병렬 처리** | 고루틴(goroutine) 기반, 경량| 스레드(thread) 기반, 상대적으로 무거움 |
| **객체지향 지원** | 제한적 지원 (클래스와 상속 없음)| 완전한 객체지향 (클래스와 상속 지원) |
| **컴파일과 실행** | 네이티브 바이너리로 컴파일, 실행 속도 빠름 | JVM에서 바이트코드 실행, 플랫폼 독립성 강함 |
| **에러 처리 방식**|  명시적 에러 반환 (`error` 타입 사용) | 예외(exception) 기반 에러 처리 |
| **문법** | 간결하고 직관적 | 풍부하지만 복잡한 문법 |

## #02. Go 언어 개발 환경 구성하기

### 1. 컴파일러 설치

(https://go.dev/dl/)[https://go.dev/dl/] 페이지에서 운영체제에 맞는 버전 설치

설치 후 명령 프롬프트에서 `go version` 명령으로 설치 여부 확인

맥의 경우 `brew install go` 명령으로도 설치 가능함

### 2. Visual Studio Code Extensions 설치

| 이름 | 개발자 |
|---|---|
| Code Runner | Jun Han |
| Go | Go Team at Google |
| Go Outliner | Go code outline explorer |

## #03. Hello Go

작업 폴더를 생성하고 1단원 폴더인 `01-HelloGo` 폴더 생성

해당 폴더 안에 `01-HelloWorld.go` 파일을 생성하고 아래의 코드를 작성

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

소스코드 작성 후 `Ctrl+Alt+N` 키로 코드 실행