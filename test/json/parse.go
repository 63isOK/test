package main

import "fmt"

// Result is parse result
type Result struct {
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
