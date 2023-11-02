package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
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

	//json := `{"name": {"first": "Konstantin", "last": "Bolotin"}, "age": 28}`

	//if !gjson.Valid(json) {
	//	panic("JSON IS NOT VALID")
	//}

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	value := gjson.Get(json, `friends.#(age>=44)#.first|@case:upper`)
	for _, last := range value.Array() {
		fmt.Println(last.String())
	}

	result, ok := gjson.Parse(json).Value().(map[string]interface{})
	if !ok {
		panic("NOT OKAY PARSING TO MAP")
	}
	fmt.Println(result)

	//fmt.Println(value.String())
	//value = gjson.Get(json, `fav*`)
	//fmt.Println(value.String())
	//fmt.Println(gjson.Parse(json).Get("name"))

}
