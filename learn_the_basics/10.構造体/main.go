package main

import "fmt"

func main() {
	// 構造体を定義
	type Vertex struct {
		x int
		y int
	}

	// 初期化
	var v1 Vertex
	v1.x = 10
	v1.y = 20
	fmt.Println(v1.x) // 10
	fmt.Println(v1.y) // 20

	// 初期化(構造体リテラル)
	v2 := Vertex{x: 10, y: 20}
	fmt.Println(v2.x) // 10
	fmt.Println(v2.y) // 20

	// 初期化(ポインタを返す)
	pv := &Vertex{x: 10, y: 20}
	fmt.Println(pv.x) // 10
	fmt.Println(pv.y) // 20
}
