package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person24 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person24) Greeting() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func task24() {
	data, err := os.ReadFile("task24.json")
	if err != nil {
		// mainなど最終的にエラーキャッチする部分で通常は使用する
		log.Fatal(err)
	}

	var p Person24
	// ポインタをunmarshalに渡す
	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Greeting())

	p.Name = "Bob"

	json, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
