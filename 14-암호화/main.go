package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	// 1. 키 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// 키 출력
	fmt.Println("Private Key:", privateKey)
	fmt.Println("Public Key:", publicKey)

	// 2. 메시지 준비 및 해시
	message := []byte("hello world")
	hash := sha256.Sum256(message)
	fmt.Println("메시지 해시:", hash)

	// 3. 서명 생성
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("서명:", signature)

	// 4. 검증
	valid := ecdsa.VerifyASN1(publicKey, hash[:], signature)
	if valid {
		fmt.Println(">>> 서명 검증 성공!")
		fmt.Println(">>> 원문 메시지:", string(message))
	} else {
		fmt.Println(">>> 서명 검증 실패!")
	}

	// 5. 변조된 메시지 테스트
	fakeMessage := []byte("hello gopher")
	fakeHash := sha256.Sum256(fakeMessage)
	fakeValid := ecdsa.VerifyASN1(publicKey, fakeHash[:], signature)
	fmt.Printf(">>> 변조 메시지 검증 결과: %v\n", fakeValid)
}
