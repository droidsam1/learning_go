package book

import "encoding/json"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b Book) toJSON() (string, error) {
	data, error := json.Marshal(b)
	if error != nil {
		return "", error
	}
	return string(data), nil
}
