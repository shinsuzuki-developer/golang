package main

import "fmt"

// 引数あり､戻り値なし
func outputNum(x int) {
	fmt.Println(x)
}

// 引数あり､戻り値あり
func calcNum1(x int, y int) int {
	return x + y
}

// 引数あり､戻り値(複数)あり
func calcNum2(x int, y int) (int, int) {
	return x + y, x - y
}

// 引数あり､戻り値(複数)あり
func calcNum3(x int, y int) (r1 int, r2 int) {
	r1 = x + y
	r2 = x - y

	//return r1, r2
	return
}

func main() {
	// 引数あり､戻り値なし
	outputNum(100) // 100

	// 引数あり､戻り値あり
	var r1 = calcNum1(10, 20)
	fmt.Println(r1) // 30

	// 引数あり､戻り値(複数)あり
	var r2, r3 = calcNum2(10, 2)
	fmt.Println(r2, r3) // 12 8

	// 引数あり､戻り値(複数)あり
	var r4, r5 = calcNum3(10, 2)
	fmt.Println(r4, r5) // 12 8

	// 関数を変数に渡して実行
	var f = func(str string) {
		fmt.Println("inner fuc:" + str)
	}

	f("execute") // inner fuc:execute

	// 関数を定義し､そのまま実行
	func(str string) {
		fmt.Println("inner fuc:" + str)
	}("exceute") // inner fuc:execute
}
