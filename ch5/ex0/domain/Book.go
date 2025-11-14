package main

import "encoding/json"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b Book) ToJson() string {
	data, _ := json.Marshal(b)
	return string(data)
}
