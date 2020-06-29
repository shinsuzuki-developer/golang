package main

import "fmt"

func main() {

	// 宣言と定義を別々に行う
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Println(a1) //[100 200]

	// 宣言と定義を同時に行う
	var a2 = [2]int{1, 2}
	fmt.Println(a2) // [1 2]

	// 宣言と定義を同時に行う､個数指定なし
	var a3 = [...]int{1, 2, 3, 4}
	fmt.Println(a3) // [1 2 3 4]

	// 参照
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	/*
		1
		2
		3
		4
	*/

	// 参照(range)
	for i, m := range a3 {
		fmt.Printf("index:%v - value:%v\n", i, m)
	}
	/*
		index:0 - value:1
		index:1 - value:2
		index:2 - value:3
		index:3 - value:4
	*/
}
