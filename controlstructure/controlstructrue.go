package main

import "fmt"

func main() {
	//학생들의 점수 목록
	students := map[string]int{
		"Alice":   85,
		"Bob":     72,
		"Charlie": 90,
		"Diana":   65,
		"Edward":  40,
	}

	//학생들의 점수를 순회하기 위한 반복문
	for name, score := range students {
		//if-else 문을 사용하여 점수에 따른 학점 계산
		var grade string
		if score >= 90 {
			grade = "A"
		} else if score >= 80 {
			grade = "B"
		} else if score >= 60 {
			grade = "C"
		} else if score >= 60 {
			grade = "D"
		} else {
			grade = "F"
		}

		//switch 문 사용하여 학점에 따른 결과 출력
		switch grade {
		case "A", "B", "C", "D":
			fmt.Printf("%s 님은 %s 학점 통과하셨습니다 (%d 점)\n", name, grade, score)
		case "F":
			fmt.Printf("%s 님은 %s 학점으로 불합격하셨습니다 (%d 점)\n", name, grade, score)
		}

	}

	fmt.Println("/n통과한 학생 목록:")
	//통과한 학생들만 출력하기 위한 반복문
	for name, score := range students {
		if score >= 60 {
			fmt.Printf("- %s (%d 점)\n", name, score)
		}
	}
}
