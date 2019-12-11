package main

import (
	"strings"
	"testing"
)

const (
	json1 = `
{
    "a":"b",
    "b":1,
    "c":{
        "d":2,
        "e":3,
        "f":"4",
        "g":[
            1,2,3,4,5
        ],
        "h":[
            "6"
        ]
    },
    "i":true,
    "j":{
        "k":null,
        "l":3.3,
        "m":[
            "a","b","c","d"
        ],
        "n":{
            "o":4.4,
            "p":"5.5"
        }
    }
}`
)

func TestParse(t *testing.T) {
	t.Run("json parse", func(t *testing.T) {
		ret := JSONParse(json1)
		checkParseResult(t, ret)
	})

	t.Run("gabs parse", func(t *testing.T) {
		ret := GabsParse(json1)
		checkParseResult(t, ret)
	})
}

func TestWrite(t *testing.T) {
	t.Run("json write", func(t *testing.T) {
		ret := JSONWrite()
		checkWriteResult(t, ret)
	})

	t.Run("gabs write", func(t *testing.T) {
		ret := GabsWrite()
		checkWriteResult(t, ret)
	})
}

func TestDParse(t *testing.T) {
	t.Run("json parse unknown type", func(t *testing.T) {
		ret := JSONDParse(json1)
		checkParseResult(t, ret)
	})

	t.Run("gabs parse unknown type", func(t *testing.T) {
		ret := GabsDParse(json1)
		checkParseResult(t, ret)
	})
}

func checkParseResult(t *testing.T, ret *Result) {
	// t.Helper()

	if ret == nil {
		t.Fatal("parse failed, result is nil")
	}

	if ret.A != "b" ||
		ret.B != 1 ||
		ret.C.D != 2 ||
		ret.C.E != 3 ||
		ret.C.F != "4" ||
		len(ret.C.G) != 5 ||
		ret.C.H[0] != "6" ||
		!ret.I ||
		ret.J.K != nil ||
		ret.J.L != 3.3 ||
		ret.J.M[2] != "c" ||
		ret.J.N.O != 4.4 ||
		ret.J.N.P != "5.5" {
		t.Fatalf("parse failed, result is %v", ret)
	}
}

func trim(s string) string {
	news := strings.ReplaceAll(s, " ", "")
	news = strings.ReplaceAll(news, "\t", "")
	news = strings.ReplaceAll(news, "\r", "")
	news = strings.ReplaceAll(news, "\n", "")
	return news
}

func checkWriteResult(t *testing.T, ret string) {
	t.Helper()

	trimJ1, trimRet := trim(json1), trim(ret)

	if trimJ1 != trimRet {
		t.Fatalf("wirte is failed, %s", ret)
	}
}
