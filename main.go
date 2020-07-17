package main

import "fmt"

const TEXT = "Text"

/*
コメント
*/
func main() {
	// 型推論で宣言
	num := 1
	str := "ABC"
	fmt.Println("Hello World",
		num,
		"aaa",
		str)

	// まとめて宣言
	var (
		a1 = "A1"
		a2 = 1
	)
	fmt.Println(a1, a2)

	// 代入2連発
	var (
		name, age = "Yamada",
			26
	)
	fmt.Println(name, age)

	// 定数
	const NAME = "Sato"
	fmt.Println(TEXT, NAME)
}
