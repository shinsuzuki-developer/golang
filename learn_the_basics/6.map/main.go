package main

import "fmt"

func main() {

	// 宣言と定義
	var m1 = map[string]int{
		"apple":  100,
		"orange": 200,
		"banana": 300,
	}

	fmt.Println(m1) // map[apple:100 banana:300 orange:200]

	// 宣言､追加
	var m2 = map[string]int{}
	m2["add-item-1"] = 100
	m2["add-item-2"] = 200
	fmt.Println(m2) // map[add-item-1:100 add-item-2:200]

	// 参照
	for key, value := range m2 {
		fmt.Println(key, value)
	}
	/*
		add-item-1 100
		add-item-2 200
	*/

	// キーの存在確認
	{
		var val, isEexist = m2["add-item-1"]
		fmt.Println(val, isEexist) // 100 true
	}

	{
		var val, isEexist = m2["nothing"]
		fmt.Println(val, isEexist) // 0 false
	}

	// makeにより初期化
	var m3 = make(map[string]int)
	m3["add-item-1"] = 100
	m3["add-item-2"] = 200
	fmt.Println(m3) // map[add-item-1:100 add-item-2:200]

}
