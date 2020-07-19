package mypkg

import "fmt"

// AnyTypeArgMethod interface{}型の引数は何でも入る
func AnyTypeArgMethod(any interface{}) {

	iVal, ok := any.(int32)
	if ok {
		fmt.Printf("int.Value=%d\n", iVal)
		return
	}
	fVal, ok := any.(float32)
	if ok {
		fmt.Printf("float32.Value=%f\n", fVal)
		return
	}
	stringable, ok := any.(ToStringable)
	if ok {
		fmt.Println("ToStringable.Value=", stringable.ToString())
		return
	}
	fmt.Println("Not suppoted type")
}

// AnyTypeArgMethodSwitch interface{}型の引数から型情報を取得してswitch分岐もできる
func AnyTypeArgMethodSwitch(any interface{}) {
	switch any.(type) {
	case bool:
	case int32:
	case float32:
	case ToStringable:
	case string:
	default:
	}
}

// any interface{}の別名
type any interface{}

// json 風のDictionary
type json map[string]any

// AnyTypeMap 型を持たないMap
func AnyTypeMap() {
	anyMap := map[any]any{
		1:     "aaa",
		"bbb": 100,
	}
	fmt.Println("--------------------")
	fmt.Println(anyMap)
	fmt.Println("--------------------")

	jsonStruct := json{
		"Name": "Date",
		"Age":  40,
		"Clothes": json{
			"Head": "Hut",
			"Eye":  "Sunglass",
			"Foot": "Leather shoes",
		},
	}
	fmt.Println(jsonStruct)
	fmt.Println(jsonStruct["Clothes"].(json)["Eye"])
	fmt.Println("--------------------")

}
