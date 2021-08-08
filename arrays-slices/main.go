package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 4}
	a[4] = 5
	fmt.Println("a : ", a)

	var b [3]int
	b[0] = 11
	b[1] = 22
	b[2] = 33
	fmt.Println("b : ", b)

	var c []int
	fmt.Println("empty c : ", c)
	c = []int{45, 63, 65}
	fmt.Println("not empty c : ", c)

	c = append(c, 22, 4, 5, 6, 78, 4)
	fmt.Println("after appending c : ", c)

	for i := 0; i < len(a); i++ {
		//fmt.Println("a[%d] : %d\n", i, a[i])
		fmt.Println("a[", i, "] is ", a[i])
	}

	for k, v := range b {
		fmt.Println("k : ", k, " v : ", v)
	}
}
