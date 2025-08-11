package main

import "fmt"

func main() {
	/** 맵 선언 및 초기화 */

	// 맵 리터럴을 이용한 초기화. 키-값 쌍을 중괄호 `{}` 안에 나열한다.
	// Java의 `HashMap<String, Integer> map = new HashMap<>();`와 유사.
	scores := map[string]int{"Alice": 90, "Bob": 85}
	fmt.Println("초기 맵:", scores)

	// `make` 함수를 이용한 맵 생성. 맵은 반드시 `make`로 초기화해야 한다.
	// `make(map[KeyType]ValueType, initialCapacity)`
	// `initialCapacity`는 선택 사항이며, 맵의 초기 크기를 미리 할당하여 성능을 최적화할 수 있다.
	// Java의 `new HashMap<>(initialCapacity);`와 유사.
	playerStats := make(map[string]int)
	fmt.Println("빈 맵 (make로 생성):", playerStats)

	/** 맵에 값 추가 및 수정 */

	// 값 추가: `map[key] = value` 형태로 추가한다. 키가 없으면 추가되고, 있으면 값이 수정된다.
	// Java의 `map.put("key", value);`와 유사.
	playerStats["Health"] = 100
	playerStats["Mana"] = 50
	fmt.Println("값 추가 후:", playerStats)

	// 값 수정: 기존 키에 새로운 값을 할당하면 값이 수정된다.
	// Java의 `map.put("key", newValue);`와 유사.
	playerStats["Health"] = 90
	fmt.Println("값 수정 후:", playerStats)

	/** 맵에서 값 가져오기 및 키 존재 여부 확인 */

	// 값 가져오기: `map[key]` 형태로 값을 가져온다. 키가 없으면 값 타입의 제로 값을 반환한다.
	// Java의 `map.get("key");`와 유사하지만, Go는 키가 없을 때 `null` 대신 제로 값을 반환한다.
	mana := playerStats["Mana"]
	fmt.Println("Mana:", mana)

	// 키 존재 여부 확인 (comma ok idiom): Go의 독특한 문법으로, 값과 함께 키의 존재 여부를 `bool` 값으로 반환한다.
	// Java의 `map.containsKey("key");`와 유사하다. Go는 값과 함께 존재 여부를 반환한다.
	hp, ok := playerStats["Health"]
	if ok { // `ok`가 `true`이면 키가 존재한다.
		fmt.Println("Health 값:", hp)
	} else {
		fmt.Println("Health 키가 존재하지 않는다.")
	}

	gold, ok := playerStats["Gold"]
	if ok {
		fmt.Println("Gold 값:", gold)
	} else {
		fmt.Println("Gold 키가 존재하지 않는다.")
	}

	/** 맵에서 값 삭제 */

	// `delete` 함수를 사용하여 키-값 쌍을 삭제한다.
	// 존재하지 않는 키를 삭제해도 런타임 에러가 발생하지 않는다.
	// Java의 `map.remove("key");`와 유사.
	delete(playerStats, "Mana")
	fmt.Println("Mana 삭제 후:", playerStats)

	delete(playerStats, "NonExistentKey") // 존재하지 않는 키 삭제 시도
	fmt.Println("존재하지 않는 키 삭제 시도 후:", playerStats)

	/** 맵의 특징 */

	// 맵은 순서가 보장되지 않는다. 순회 시마다 순서가 달라질 수 있다.
	// Java의 `HashMap`도 순서가 보장되지 않는다. 순서가 필요한 경우 `LinkedHashMap`을 사용한다.
	// 맵의 제로 값은 `nil`이다. `nil` 맵은 요소를 추가할 수 없으며, 사용하려면 `make`로 초기화해야 한다.
	// Java의 `null` 맵에 `put` 시 `NullPointerException` 발생과 유사하다.
	var nilMap map[string]int
	fmt.Printf("nilMap: %v, nil 여부: %t\n", nilMap, nilMap == nil)
	// nilMap["test"] = 10 // 런타임 패닉 발생: assignment to entry in nil map
}
