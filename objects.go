package main

import "fmt"

func main() {
	me["name"] = "Kimi"
	me["age"] = 1
	me["hair color"] = "none"

	for key := range me {
		fmt.Println(key)
	}
}

// 1. Using map index notation, modify the existing values to replace the information with your own

var me = map[string]any{
	"name":       "Jake Overall",
	"age":        29,
	"hair color": "brown",
}

// 2. Iterate over the map to print the keys
