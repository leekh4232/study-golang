package main

import "fmt"

// GameCharacter 구조체를 정의함.
// Java의 클래스와 유사하게 관련 데이터(필드)를 묶는 역할을 함.
// 하지만 Java 클래스와 달리 메서드를 포함하지 않고 순수하게 데이터만 정의함.
type GameCharacter struct {
    Name  string // 캐릭터의 이름을 저장하는 필드
    Level int    // 캐릭터의 레벨을 저장하는 필드
    HP    int    // 캐릭터의 체력을 저장하는 필드
    MP    int    // 캐릭터의 마나를 저장하는 필드
}

// 프로그램의 시작점인 main 함수임.
func main() {
    // GameCharacter 구조체의 인스턴스(객체)를 생성함.
    // var 키워드를 사용하여 'character1'이라는 이름의 변수를 선언하고, GameCharacter 타입으로 지정함.
    var character1 GameCharacter

    // 필드에 값을 할당함. .(점) 연산자를 사용하여 각 필드에 접근함.
    character1.Name = "Hossam"
    character1.Level = 99
    character1.HP = 8000
    character1.MP = 12000

    // 구조체 인스턴스를 초기화하는 또 다른 방법 (리터럴 초기화)
    // 선언과 동시에 중괄호 {} 안에 필드와 값을 직접 명시하여 인스턴스를 생성함.
    character2 := GameCharacter{
        Name:  "Alex",
        Level: 50,
        HP:    5000,
        MP:    7000,
    }

    // 생성된 구조체 인스턴스의 정보를 출력함.
    // Println 함수는 구조체를 출력할 때 필드와 값을 보기 좋게 표시해줌.
    fmt.Println("캐릭터 1 정보:", character1)
    fmt.Println("캐릭터 2 정보:", character2)

    // 특정 필드에 접근하여 값을 개별적으로 출력할 수도 있음.
    fmt.Printf("캐릭터 2의 이름: %s, 레벨: %d\n", character2.Name, character2.Level)
}