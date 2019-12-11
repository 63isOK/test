package main

import (
	"encoding/json"
	"fmt"
)

type structc struct {
	D int      `json:"d"`
	E int      `json:"e"`
	F string   `json:"f"`
	G []int    `json:"g"`
	H []string `json:"h"`
}

type structn struct {
	O float64 `json:"o"`
	P string  `json:"p"`
}

type structj struct {
	K *int     `json:"k"`
	L float64  `json:"l"`
	M []string `json:"m"`
	N structn  `json:"n"`
}

// Result is parse result
type Result struct {
	A string  `json:"a"`
	B int     `json:"b"`
	C structc `json:"c"`
	I bool    `json:"i"`
	J structj `json:"j"`
}

// JSONParse parse static json with encoding/json
func JSONParse(s string) *Result {
	var ret Result

	err := json.Unmarshal([]byte(s), &ret)

	if err != nil {
		fmt.Printf("encoding/json parse failed: %s\n %s\n", s, err.Error())
		return nil
	}

	return &ret
}

// GabsParse parse static json with gabs
func GabsParse(s string) *Result {
	return nil
}

// JSONWrite write static json with encoding/json
func JSONWrite() string {
	v := Result{
		A: "b",
		B: 1,
		C: structc{
			D: 2,
			E: 3,
			F: "4",
			G: []int{1, 2, 3, 4, 5},
			H: []string{"6"},
		},
		I: true,
		J: structj{
			K: nil,
			L: 3.3,
			M: []string{"a", "b", "c", "d"},
			N: structn{
				O: 4.4,
				P: "5.5",
			},
		},
	}

	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("encode failed ", err.Error())
		return ""
	}

	return string(b)
}

// GabsWrite write static json with gabs
func GabsWrite() string {
	return ""
}

// JSONDParse parse dynamic json with encoding/json
func JSONDParse(s string) *Result {
	return nil
}

// GabsDParse parse dynamic json with gabs
func GabsDParse(s string) *Result {
	return nil
}

func main() {
	fmt.Println("反序列化时，结构体的字段要是大写才能参与解码过程")
}
