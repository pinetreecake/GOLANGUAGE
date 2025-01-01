package main

import "fmt"

func main() {
	//-----변수(Variable)------

	//1.초기화 없이 변수 선언 (기본값 자동 할당)
	var age int                   //int 타입 기본값 0
	fmt.Println("기본 age 값:", age) //출력: 기본 age 값:0

	//2. 초기값과 함께 변수 선언
	var name string = "Alice"
	fmt.Println("이름:", name) //출력: 이름: Alice

	//3. 타입 추론 (타입 자동 감지)
	city := "New York"       // var city string = "New York"와 동일
	fmt.Println("도시:", city) //출력: 도시: New York

	//4. 여러 변수 동시 선언
	var height, weight int = 170, 65
	fmt.Println("키:", height, "cm,몸무게:", weight, "kg") //출력: 키" 170cm, 몸무게: 65kg

	//5. 변수 값 업데이트
	age = 25
	fmt.Println("업데이트된 age 값:", age) //출력: 업데이트 된 age 값: 25

	//--상수 (Constants)--

	//상수는 선언 이후 변경할 수 없음
	const pi float64 = 3.14159
	fmt.Println("Pi 값:", pi) //츌력: Pi 값: 3.14159

	//여러 상수 선언
	const (
		EarthGravity = 9.8 //지구의 중력 가속도
		MoonGravity  = 1.6 //달의 중력 가속도
	)
	fmt.Println("지구 중력:", EarthGravity, "m/s2") //출력: 지구 중력: 9.8 m/s2
	fmt.Println("달 중력:", MoonGravity, "m/s2")   //출력: 달 중력: 1.6 m/s2

	//변수와 상수 활용한 예제
	circumference := 2 * pi * float64(height)             //상수 'pi' 사용
	fmt.Println("반지름이", height, "인원의 둘레:", circumference) //height에 따라 출력 값 변동
}
