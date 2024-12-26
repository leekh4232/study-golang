package main

import "fmt"

func main() {
	// 1. 배열 (array): 고정 크기의 유닛 체력 정보
	var unitHP [3]int = [3]int{40, 35, 150} // 마린, 저글링, 시즈탱크 체력

	fmt.Println("유닛 체력 (배열):")
	fmt.Println("마린 체력:", unitHP[0])
	fmt.Println("저글링 체력:", unitHP[1])
	fmt.Println("시즈탱크 체력:", unitHP[2])

	// 2. 슬라이스 (slice): 유닛 이름 리스트
	var unitNames []string = []string{"Marine", "Zergling", "Siege Tank"}

	fmt.Println("\n유닛 이름 (슬라이스):")
	fmt.Println("첫 번째 유닛:", unitNames[0])
	fmt.Println("두 번째 유닛:", unitNames[1])
	fmt.Println("세 번째 유닛:", unitNames[2])

	// 3. 맵 (map): 유닛 공격력 매핑
	var unitAttack map[string]int = map[string]int{
		"Marine":     6,
		"Zergling":   5,
		"Siege Tank": 70,
	}

	fmt.Println("\n유닛 공격력 (맵):")
	fmt.Println("마린 공격력:", unitAttack["Marine"])
	fmt.Println("저글링 공격력:", unitAttack["Zergling"])
	fmt.Println("시즈탱크 공격력:", unitAttack["Siege Tank"])

	// 4. 구조체 (struct): 유닛 상세 정보
	type Unit struct {
		Name     string
		HP       int
		Attack   int
		IsFlying bool
	}

	var marine Unit = Unit{Name: "Marine", HP: 40, Attack: 6, IsFlying: false}
	var zergling Unit = Unit{Name: "Zergling", HP: 35, Attack: 5, IsFlying: false}
	var siegeTank Unit = Unit{Name: "Siege Tank", HP: 150, Attack: 70, IsFlying: false}

	fmt.Println("\n유닛 상세 정보 (구조체):")
	fmt.Println("마린 이름:", marine.Name, "체력:", marine.HP, "공격력:", marine.Attack, "공중 여부:", marine.IsFlying)
	fmt.Println("저글링 이름:", zergling.Name, "체력:", zergling.HP, "공격력:", zergling.Attack, "공중 여부:", zergling.IsFlying)
	fmt.Println("시즈탱크 이름:", siegeTank.Name, "체력:", siegeTank.HP, "공격력:", siegeTank.Attack, "공중 여부:", siegeTank.IsFlying)

	// 5. 채널 (chan): 유닛 명령 전달
	var unitCommands chan string = make(chan string, 3)
	unitCommands <- "Move to Point A"
	unitCommands <- "Attack Enemy"
	unitCommands <- "Hold Position"

	fmt.Println("\n유닛 명령 (채널):")
	fmt.Println("첫 번째 명령:", <-unitCommands)
	fmt.Println("두 번째 명령:", <-unitCommands)
	fmt.Println("세 번째 명령:", <-unitCommands)
}
