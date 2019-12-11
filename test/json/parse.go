package main

import (
	"encoding/json"
	"fmt"
)

type structc struct {
	D, E int
	F    string
	G    []int
	H    []string
}

type structn struct {
	O float64
	P string
}

type structj struct {
	K *int
	L float64
	M []string
	N structn
}

// Result is parse result
type Result struct {
	A string `json:"a"`
	B int
	C structc
	I bool
	J structj
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
	return ""
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
