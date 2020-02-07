package utils

import (
	"fmt"
	"testing"

	"github.com/fatih/structs"
)

func TestConvertStructToMap(t *testing.T) {
	type Server struct {
		Name    string `json:"name"`
		ID      int    `json:"id"`
		Enabled bool   `json:"enabled"`
		User    string `json:"user"`
	}

	server := &Server{
		Name: "gopher",
		ID:   123456,
		User: "",
	}

	var m = make(map[string]interface{}, 2)
	s := structs.New(server)
	for _, f := range s.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())
		fmt.Println(f.Value())
		if f.IsExported() {
			if !f.IsZero() {
				m[f.Name()] = f.Value()
			}
			fmt.Printf("value   : %+v\n", f.Value())
			fmt.Printf("is zero : %+v\n", f.IsZero())
		}
	}
	fmt.Println(m)
}

func TestConvertStructToMap2(t *testing.T) {

}
