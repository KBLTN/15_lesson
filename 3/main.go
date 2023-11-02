package main

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func main() {

	json := `{
	"name": {"first": "Olesia", "last": "Bolotin"},
	"age" : 26,
	"children": ["Andrey", "Yana", "Victoria"],
	"fav.movie": "That's My Boy",
	"friends": [
	{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]}
	{"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb", "tw"]}
	{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	}`

	val, _ := sjson.Delete(json, "friends.1")
	fmt.Println(val)

	//value, _ := sjson.Set(json, "new_object", map[string]interface{}{"hi": "Kotya"})
	//fmt.Println(value)
}
