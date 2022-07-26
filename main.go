package main

import (
	"PPJson/json"
	"fmt"
)

type test struct {
	A string                 `json:"AAA"`
	B int                    `json:"Test"`
	C bool                   `json:"CTests"`
	T tt                     `json:"TTT"`
	M map[string]interface{} `json:"Mmap"`
}

type tt struct {
	A string
	B []int
}

func main() {

	t := test{
		A: "!@34",
		B: 123,
		C: true,
		T: struct {
			A string
			B []int
		}{A: "Hao_pp", B: []int{1, 2, 3, 4, 5, 6}},
		M: map[string]interface{}{
			"Test": "Ttt",
			"123":  111,
			"qq":   1446785380,
		},
	}

	fmt.Println("原结构体:")
	fmt.Println(t, "\n")

	fmt.Println("JSON序列化:")
	ans, _ := json.JSONMarshal(t)

	fmt.Println(ans, "\n")

	var t1 test

	fmt.Println("JSON反序列化:")
	json.JSONUnMarshal(&t1, ans)

	fmt.Println(t1)

}
