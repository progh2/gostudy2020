package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("안녕, 세상아?!")

	// 처음 정의 및 할당 할 때는 := 사용. 그 후 할당 때는 = 사용.
	sum := 0
	// for 초기화 조건; 조건식; 후속작업 정의
	// go에는 while이 없다.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2, i2 := 0, 0
	for i2 < 10 {
		sum2 += i2
		i2++
	}
	fmt.Println(sum2)

	sum3, i3 := 0, 0
	// for 문에 조건식 생략
	for{
		if i3>= 10 {
			break
		}
		sum3 += i3
		i3++
	}
	fmt.Println(sum3)
}
