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
	t.Helper()

	if ret == nil {
		t.Fatal("parse failed, result is nil")
	}

	if ret.a != "b" ||
		ret.b != 1 ||
		ret.c.d != 2 ||
		ret.c.e != 3 ||
		ret.c.f != "4" ||
		len(ret.c.g) != 5 ||
		ret.c.h[0] != "6" ||
		!ret.i ||
		ret.j.k != nil ||
		ret.j.l != 3.3 ||
		ret.j.m[2] != "c" ||
		ret.j.n.o != 4.4 ||
		ret.j.n.p != "5.5" {
		t.Fatal("parse failed, result is nil")
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
