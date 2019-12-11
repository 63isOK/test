package main

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
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

// // GabsParse parse static json with gabs
// func GabsParse(s string) *Result {
//   jsonParsed, err := gabs.ParseJSON([]byte(s))
//   if err != nil {
//     fmt.Printf("Jeffail/gabs parse failed: %s\n %s\n", s, err.Error())
//     return nil
//   }
//
//   // // 这儿的类型断言是不通过的，因为gabs.Container里的类型是interface{}
//   // // 数据类型的转换，只能通过遍历gabs.container + 给结构体赋值的方式来实现
//   // v, ok := jsonParsed.Data().(Result)
//   // if !ok {
//   //   fmt.Println(jsonParsed.String())
//   //   return nil
//   // }
//   //
//   // return &v
//
//   // 将已知格式的json反序列化到指定数据类型，直接使用encoding/json即可
//
//   return nil
// }

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
	obj := gabs.New()
	obj.SetP("b", "a")
	obj.SetP(1, "b")
	obj.SetP(2, "c.d")
	obj.SetP(3, "c.e")
	obj.SetP("4", "c.f")
	obj.ArrayP("c.g")
	obj.ArrayAppendP(1, "c.g")
	obj.ArrayAppendP(2, "c.g")
	obj.ArrayAppendP(3, "c.g")
	obj.ArrayAppendP(4, "c.g")
	obj.ArrayAppendP(5, "c.g")
	obj.ArrayP("c.h")
	obj.ArrayAppendP("6", "c.h")
	obj.SetP(true, "i")
	obj.SetP(nil, "j.k")
	obj.SetP(3.3, "j.l")
	obj.ArrayP("j.m")
	obj.ArrayAppendP("a", "j.m")
	obj.ArrayAppendP("b", "j.m")
	obj.ArrayAppendP("c", "j.m")
	obj.ArrayAppendP("d", "j.m")
	obj.SetP(4.4, "j.n.o")
	obj.SetP("5.5", "j.n.p")

	return obj.String()
}

// // JSONDParse parse dynamic json with encoding/json
// func JSONDParse(s string) *Result {
//   var v interface{}
//   err := json.Unmarshal([]byte(s), &v)
//   if err != nil {
//     fmt.Println("encoding/json parse dynamic json failed ", err.Error())
//     return nil
//   }
//
//   // 解析未知格式的json，或是动态组Go类型去序列化成json
//   // 都是蛮复杂的，这里只是为了了解encoding/json的能力，就不进一步深入了
//   // 直接上gabs
// }
//
// // GabsDParse parse dynamic json with gabs
// func GabsDParse(s string) *Result {
//   return nil
// }

func main() {
	fmt.Println("反序列化时，结构体的字段要是大写才能参与解码过程")
}
