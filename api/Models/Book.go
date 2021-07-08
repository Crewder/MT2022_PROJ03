package Models

import (
	"encoding/json"
	"reflect"
)

type Book struct {
	Author   string `json:"author,omitempty"`
	Abstract string `json:"abstract,omitempty"`
	Title    string `json:"title,omitempty"`
}

func HydrateBook(input Book) ([]byte, error) {
	model := &Book{
		Author:   input.Author,
		Abstract: input.Abstract,
		Title:    input.Title,
	}
	return json.Marshal(model)
}

func GetFields() []string {
	var input Book
	var fields []string

	v := reflect.ValueOf(input)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	return fields
}
