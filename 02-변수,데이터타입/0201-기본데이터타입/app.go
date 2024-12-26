// 동일 폴더 내의 같은 패키지는 하나의 실행 파일로 묶인다.
// 이 때, main함수는 하나만 존재해야 하므로 폴더를 구분하여 main의 중복 선언 문제를 해결한다.
package main

import "fmt"

func main() {
	// 1. 유닛 정보 정의
	var unitName string = "Marine" // 유닛 이름
	var unitHP int = 40            // 유닛의 체력
	var unitAttack uint = 6        // 유닛의 공격력
	var unitArmor uint = 0         // 유닛의 방어력
	var attackSpeed float32 = 0.86 // 유닛의 공격 속도
	var isFlying bool = false      // 공중 유닛 여부

	// 2. 유닛 정보 출력
	fmt.Println("유닛 정보:")
	fmt.Println("이름:", unitName)
	fmt.Println("체력:", unitHP)
	fmt.Println("공격력:", unitAttack)
	fmt.Println("방어력:", unitArmor)
	fmt.Println("공격 속도:", attackSpeed)
	fmt.Println("공중 유닛 여부:", isFlying)

	// 3. 적 유닛 정보 정의 및 출력
	var enemyName string = "Zergling"
	var enemyHP int = 35
	var enemyArmor uint = 0

	fmt.Println("\n적 유닛 정보:")
	fmt.Println("이름:", enemyName)
	fmt.Println("체력:", enemyHP)
	fmt.Println("방어력:", enemyArmor)

	// 4. 추가 유닛 리스트 정의
	var unit1 string = "Marine"
	var unit2 string = "Firebat"
	var unit3 string = "Medic"

	fmt.Println("\n유닛 리스트:")
	fmt.Println("유닛 1:", unit1)
	fmt.Println("유닛 2:", unit2)
	fmt.Println("유닛 3:", unit3)
}
