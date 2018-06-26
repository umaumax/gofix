package main

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	filetype = "go"
	type set struct {
		in, out string
	}
	//	if などの短すぎる単語は難しい "i" や "i,"
	list := []set{
		{"retun", "return"},
		{"ret", "ret"},
		{"retun", "return"},
		{"fro", "for"},
		{"rnge", "range"},
		{"func", "func"},
		{"import", "import"},
		{"type", "type"},
		{"struct", "struct"},
		{"strcut", "struct"},
		{"tpe sample stcut", "type sample struct"},
		{"fro i := 0; i < 10; i++ {", "for i := 0; i < 10; i++ {"},
		{"fo i, v := rnge list {", "for i, v := range list {"},
		{"n**2", "n*n"},
		{"(r2-r1)**2", "(r2-r1)*(r2-r1)"},
	}
	cnt := 0
	for i, v := range list {
		result := gofix(filetype, v.in)
		if result != v.out {
			fmt.Println("failure", i, "\n", "in :", v.in, "\n", "out:", result, "\n", "req:", v.out)
		} else {
			cnt++
		}
	}
	got := cnt
	want := len(list)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
