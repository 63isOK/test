package main

import "testing"

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
        ]
    }
}`
)

func TestParse(t *testing.T) {
	t.Run("json parse", func(t *testing.T) {
		ret := JSONParse(json1)
		checkParseResult(ret)
	})

	t.Run("gabs parse", func(t *testing.T) {
		ret := GabsParse(json1)
		checkParseResult(ret)
	})
}

func TestWrite(t *testing.T) {
	t.Run("json write", func(t *testing.T) {
		ret := JSONWrite()
		checkWriteResult(ret)
	})

	t.Run("gabs write", func(t *testing.T) {
		ret := GabsWrite()
		checkWriteResult(ret)
	})
}

func TestDParse(t *testing.T) {
	t.Run("json parse unknown type", func(t *testing.T) {
		ret := JSONDParse(json1)
		checkParseResult(ret)
	})

	t.Run("gabs parse unknown type", func(t *testing.T) {
		ret := GabsDParse(json1)
		checkParseResult(ret)
	})
}

func checkParseResult(ret *Result) {
}

func checkWriteResult(ret string) {

}
