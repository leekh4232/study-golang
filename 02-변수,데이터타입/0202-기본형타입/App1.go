package main

import "fmt"

func main() {
	// 1. 유닛 기본 속성 (정수형 데이터 타입)
	var marineAttack int8 = 6      // 마린 공격력 (int8)
	var marineHP int16 = 40        // 마린 체력 (int16)
	var zerglingAttack int8 = 5    // 저글링 공격력 (int8)
	var zerglingHP int16 = 35      // 저글링 체력 (int16)
	var siegeTankAttack int32 = 70 // 시즈탱크 공격력 (int32)
	var siegeTankHP int64 = 150    // 시즈탱크 체력 (int64)

	fmt.Println("유닛 기본 속성:")
	fmt.Println("마린 - 공격력:", marineAttack, "체력:", marineHP)
	fmt.Println("저글링 - 공격력:", zerglingAttack, "체력:", zerglingHP)
	fmt.Println("시즈탱크 - 공격력:", siegeTankAttack, "체력:", siegeTankHP)

	// 2. 방어력 (부호 없는 정수형 데이터 타입)
	var marineArmor uint8 = 0     // 마린 방어력
	var zerglingArmor uint8 = 0   // 저글링 방어력
	var siegeTankArmor uint16 = 2 // 시즈탱크 방어력

	fmt.Println("\n유닛 방어력:")
	fmt.Println("마린 - 방어력:", marineArmor)
	fmt.Println("저글링 - 방어력:", zerglingArmor)
	fmt.Println("시즈탱크 - 방어력:", siegeTankArmor)

	// 3. 복소수를 이용한 위치 데이터
	var marinePosition complex64 = 5 + 2i      // 마린의 위치
	var zerglingPosition complex64 = 3 + 4i    // 저글링의 위치
	var siegeTankPosition complex128 = 7 + 10i // 시즈탱크의 위치

	fmt.Println("\n유닛 위치:")
	fmt.Println("마린 위치:", marinePosition)
	fmt.Println("저글링 위치:", zerglingPosition)
	fmt.Println("시즈탱크 위치:", siegeTankPosition)

	// 4. 유니코드 아이콘 추가 (문자 데이터 타입)
	// [windows] WinKey + . (마침표) : 이모지 입력
	// [Mac]  Control + Command + 스페이스바
	// https://symbl.cc/kr/unicode-table/ 에서 복사
	var marineIcon rune = '⚔'    // 마린 아이콘 (검 모양)
	var zerglingIcon rune = '🐜'  // 저글링 아이콘 (개미 모양)
	var siegeTankIcon rune = '🚜' // 시즈탱크 아이콘 (탱크 상징)

	fmt.Println("\n유닛 아이콘:")
	fmt.Println("마린 아이콘:", string(marineIcon))
	fmt.Println("저글링 아이콘:", string(zerglingIcon))
	fmt.Println("시즈탱크 아이콘:", string(siegeTankIcon))

	// 5. 포인터 크기 데이터 타입
	var mapIdentifier uintptr = 12345678 // 맵의 고유 식별자

	fmt.Println("\n맵 정보:")
	fmt.Println("맵 식별자:", mapIdentifier)
}
