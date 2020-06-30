package main

import "fmt"

func main() {
	// 宣言と定義
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1)      // [1 2 3 4 5]
	fmt.Println(s1[2:4]) // [3 4]
	fmt.Println(s1[3:])  // [4 5]
	fmt.Println(s1[:3])  // [1 2 3]
	fmt.Println(s1[:])   // [1 2 3 4 5]

	// makeを使用した初期化
	var s2 = make([]int, 5, 5)
	fmt.Println(s2)

	// 宣言､項目を追加
	var s3 []int
	s3 = append(s3, 1)
	s3 = append(s3, 2)
	fmt.Println(s3) // [1 2]

	// 参照
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	/*
		1
		2
	*/
}
