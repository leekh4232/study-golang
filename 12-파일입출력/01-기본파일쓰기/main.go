package main

import (
	"fmt"
	"os" // 운영체제와의 인터페이스를 제공하는 패키지
)

func main() {
	// 1. 파일 생성: os.Create()는 파일을 생성하거나 기존 파일을 덮어씀
	//    Java의 new FileOutputStream("example.txt")와 유사함
	file, err := os.Create("../example.txt")
	if err != nil {
		fmt.Println("파일 생성 실패:", err)
		return // Java의 throw new IOException()과 유사한 에러 처리
	}
	// 2. defer를 사용한 리소스 관리: Java의 try-with-resources와 유사
	//    함수가 종료될 때 자동으로 파일이 닫힘
	defer file.Close()

	// 3. 파일에 쓸 데이터 준비: Go는 바이트 슬라이스를 사용
	//    Java의 byte[] 배열과 유사하지만 더 유연함
	data := []byte("Go 언어 파일 입출력 실습\n첫 번째 줄입니다.\n두 번째 줄입니다.")

	// 4. 파일에 데이터 쓰기: Write 메서드는 바이트 슬라이스를 받음
	//    Java의 OutputStream.write(byte[])와 동일한 역할
	bytesWritten, err := file.Write(data)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
		return
	}

	// 5. 결과 출력 및 동기화
	fmt.Printf("파일에 %d 바이트를 성공적으로 썼습니다.\n", bytesWritten)

	// 6. 파일 버퍼를 디스크에 강제로 쓰기 (Java의 flush()와 유사)
	err = file.Sync()
	if err != nil {
		fmt.Println("파일 동기화 실패:", err)
		return
	}

	fmt.Println("파일 쓰기 완료")
}
