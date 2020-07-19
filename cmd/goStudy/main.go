package main

import "fmt"
import "mypkg"

// Text は動作確認用の定数です
const Text = "Text"

/*
コメント
*/
func main() {
	declareVarStudy()

	typeStudy()

	mapAndSwitchStudy()

	person PersonModel*
	person.SetPerson("Tomizawa", 10338)
	toStringable ToStringable* = person
	fmt.Println("--------------------")
	fmt.Println(toStringable.ToString())
	fmt.Println("--------------------")
}

func declareVarStudy() {
	// 型推論で宣言
	num := 1
	str := "ABC"
	fmt.Println("--------------------")
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
	const (
		LName = "Sato"
		FName = "Taro"
	)
	fmt.Println(Text, LName, FName)
	fmt.Println("--------------------")
}

// UtcTime is Aliase of string
type UtcTime string

//JstTime is Aliase of string
type JstTime string

func typeStudy() {
	var (
		b1           = true
		u1   uint8   = 0x01
		u2   uint8   = 0x02                 // bit演算可能(u1 | u2 = 0x03)
		f1           = 1.234567890123456789 // 17桁保持 (1.2345678901234567)
		f2   float32 = 1.234567890123456789 // 8桁保持 (1.2345679)
		f3   float64 = 1.234567890123456789 // 17桁保持 (1.2345678901234567)
		r1   rune    = 1000                 // int32と同義らしい
		utc1 UtcTime = "12:34:56"
		jst1 JstTime = "78:90:12"
	)
	// 元をただせばstringだが、別名で定義すればエラーになるのでタイプセーフっぽく使える
	// UtcTime = JstTime

	fmt.Println("--------------------")
	fmt.Println("bool=", b1,
		",bit=", u1|u2,
		",float1=", f1,
		",float2=", f2,
		",float3=", f3,
		",rune1=", r1,
		",UtcTime=", utc1,
		",JstTime=", jst1)

	// Castすればイケる
	utc1 = UtcTime(jst1)

	fmt.Println("--------------------")
	fmt.Println("UtcTime=", utc1)

	var (
		a1 = [3]float32{}
		a2 = []float32{1, 2, 3}    // 要素数可変、パフォーマンス若干落ちる
		a3 = [...]float32{4, 5, 6} // 要素数固定、パフォーマンス維持
	)
	a2 = append(a2, 7)

	fmt.Println("--------------------")
	fmt.Println("a1=", a1,
		",a2=", a2,
		",a3=", a3)

	/* 要素数可変だが、事前にメモリを確保しておくと、
	容量をオーバーするまでは再確保のオーバーヘッドがなくなる */
	buff := make([]byte, 0, 1024)
	buff = append(buff, 0x08)
	fmt.Println("buff=", buff)

	fmt.Println("--------------------")
}

func mapAndSwitchStudy() {
	fmt.Println("--------------------")
	m1 := map[string]int{
		"tanaka": 60,
		"sato":   75,
	}
	m1["ito"] = 80
	_, ok := m1["tanaka"]
	if ok {
		delete(m1, "tanaka")
	}
	m1Len := len(m1)
	m1["len"] = m1Len
	fmt.Println("m1=", m1)

	for key, value := range m1 {
		switch key {
		case "ito":
			fmt.Println("key=", key, ",value=", value)
		case "len":
			fmt.Println("Len")
			// 下のラベルに滑り落ちたければ"fallofthrough"
			fallthrough
		default:
			fmt.Println("Default")
		}
		// if文っぽくも使えるらしい
		switch {
		case key == "ito" && value == 80:
			// おもしろいね
		}
	}
	fmt.Println("--------------------")
}
