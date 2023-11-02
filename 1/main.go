package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocker bool   `json:"is_blocker"`
	Books     []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func main() {
	byt := []byte(`{"name":"Konstantin","age":28,"is_blocker":true,"books":[{"name":"BookName","year":1991},{"name":"BookName2","year":2022}]}`)
	var dat User
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat.Books[1].Name)
}

func serialize() {
	var books []Book
	book1 := Book{
		Name: "BookName",
		Year: 1991,
	}
	book2 := Book{
		Name: "BookName2",
		Year: 2022,
	}

	books = append(books, book1, book2)
	sv := User{
		"Konstantin",
		28,
		true,
		books,
	}
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
}
