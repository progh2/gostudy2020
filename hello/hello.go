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

}
