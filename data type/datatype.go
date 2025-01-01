package main

import "fmt"

func main() {
	//--Integer(정수형)--
	//int 기본형 정수형 타입이며, 32비트 또는 64비트로 플랫폼에 따라 크기가 달라질 수 있음
	var age int = 25
	fmt.Println("정수 age 값:", age) // 츌룍: 정수 age 값: 25

	//uint는 부호가 없는 정수형 (0 이상의 값만 가질 수 있음)
	var positiveNumber uint = 42
	fmt.Println("양의 정수 positiveNumber 값:", positiveNumber) //출력: 양의 정수 positiveNumber 값:42

	//------Floating Point (부동 소수점)---
	//float32와 float64 소수점 값 표현함
	var height float32 = 170.5
	fmt.Println("부동 소수점 height 값:", height) //출력: 부동 소수점 height 값: 170.5

	var pi float64 = 3.141592653589793
	fmt.Println("부동 소수점 pi 값:", pi) //출력: 부동 소수점 pi 값: 3.141592653589793

	//--String (문자열)--
	//문자열 텍스트 데이터 나타냅니다.
	var firstName string = "Alice"
	var lastName string = "Johnson"
	fullName := firstName + " " + lastName //문자열 연결
	fmt.Println("전체 이름:", fullName)        //출력 전체 이름:Alice Johnson

	//문자열 길이를 계산할 수 있음
	fmt.Println("이름 길이:", len(fullName)) //출력: 이름 길이: 12

	//----Boolean(불리언)--
	//불리언은 참(true) 또는 거짓 (false) 을 표현함
	var isStudent bool = true
	fmt.Println("학생 여부:", isStudent) //출력: 학생 여부:true

	//논리 연산자를 활용한 예
	var isAdult bool = age >= 18
	fmt.Println("성인 여부:", isAdult) //출력: 성인 여부:true

	// ----Example: 다양한 데이터 타입 활용한 예제----
	fmt.Printf("안녕하세요 제 이름은 %s이고, 나이는 %d세입니다.\n", fullName, age)
	fmt.Printf("저의 키는 %t\n", isAdult)
	fmt.Printf("저의 성인인가요? %t\n", isAdult) //출력: 저는 성인인가요? true
}
