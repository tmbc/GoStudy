package mypkg

import "fmt"

// ToStringable はToString()を提供するインターフェイスです
type ToStringable interface {
	ToString() string
}

// Print メソッドのデフォルトメソッド
func Print(stringable ToStringable) {
	fmt.Println(stringable.ToString())
}
