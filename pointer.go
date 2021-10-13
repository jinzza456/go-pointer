package main

import "fmt"

func main() {
	var a int
	var b int
	var p *int // *로 포인터 변수를 선언

	a = 3
	b = 2
	p = &a // &로 주소값을 받는다.

	fmt.Println(a)
	fmt.Println(p)  // p 는 a의 주소를 받았기 때문에 주소가 나옴
	fmt.Println(*p) // *p를 하면 p가 받은 주소에 저장된 값이 나옴

	p = &b // b의 주소를 받음

	fmt.Println(b)
	fmt.Println(p)  // p 는 주소를 받았기 때문에 주소가 나옴
	fmt.Println(*p) // *p를 하면 p가 받은 주소에 저장된 값이 나옴
}
