package main

import "fmt"

type structc struct {
	d, e int
	f    string
	g    []int
	h    []string
}

type structn struct {
	o float64
	p string
}

type structj struct {
	k *int
	l float64
	m []string
	n structn
}

// Result is parse result
type Result struct {
	a string
	b int
	c structc
	i bool
	j structj
}

// JSONParse parse static json with encoding/json
func JSONParse(s string) *Result {
	return nil
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
	fmt.Println("vim-go")
}
