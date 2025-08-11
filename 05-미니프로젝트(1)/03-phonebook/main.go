package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 전화번호부를 저장할 맵 (이름 -> 전화번호)
var phoneBook = make(map[string]string)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 무한 루프를 통해 메뉴를 계속 표시 (Java의 `while(true)`와 유사)
	for {
		fmt.Println("\n--- 전화번호부 ---")
		fmt.Println("1. 연락처 추가")
		fmt.Println("2. 연락처 조회")
		fmt.Println("3. 연락처 삭제")
		fmt.Println("4. 모든 연락처 보기")
		fmt.Println("5. 종료")
		fmt.Print("메뉴를 선택하세요: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Go의 switch 문은 조건식 없이도 사용 가능하며, 여러 case를 묶을 수 있다.
		switch choice {
		case "1":
			addContact(reader)
		case "2":
			viewContact(reader)
		case "3":
			deleteContact(reader)
		case "4":
			listAllContacts()
		case "5":
			fmt.Println("전화번호부를 종료합니다.")
			return // main 함수 종료
		default:
			fmt.Println("유효하지 않은 메뉴입니다. 다시 선택해주세요.")
		}
	}
}

// 연락처 추가 함수
func addContact(reader *bufio.Reader) {
	fmt.Print("이름을 입력하세요: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("전화번호를 입력하세요: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	phoneBook[name] = phone // 맵에 추가 또는 업데이트
	fmt.Printf("'%s' 연락처가 추가되었습니다.\n", name)
}

// 연락처 조회 함수
func viewContact(reader *bufio.Reader) {
	fmt.Print("조회할 이름을 입력하세요: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	phone, ok := phoneBook[name] // 맵에서 값 가져오기 및 존재 여부 확인
	if ok {
		fmt.Printf("이름: %s, 전화번호: %s\n", name, phone)
	} else {
		fmt.Printf("'%s' 연락처를 찾을 수 없습니다.\n", name)
	}
}

// 연락처 삭제 함수
func deleteContact(reader *bufio.Reader) {
	fmt.Print("삭제할 이름을 입력하세요: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, ok := phoneBook[name] // 존재 여부 확인
	if ok {
		delete(phoneBook, name) // 맵에서 삭제
		fmt.Printf("'%s' 연락처가 삭제되었습니다.\n", name)
	} else {
		fmt.Printf("'%s' 연락처를 찾을 수 없습니다.\n", name)
	}
}

// 모든 연락처 보기 함수
func listAllContacts() {
	if len(phoneBook) == 0 {
		fmt.Println("전화번호부가 비어 있습니다.")
		return
	}
	fmt.Println("\n--- 모든 연락처 ---")
	// 맵 순회 (순서 보장 안됨)
	for name, phone := range phoneBook {
		fmt.Printf("이름: %s, 전화번호: %s\n", name, phone)
	}
}
