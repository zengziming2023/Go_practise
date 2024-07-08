package main

import (
	"fmt"
	ss "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", ss.Contains("test", "t"))
	p("Count: ", ss.Count("test", "t"))
	p("HasPrefix:", ss.HasPrefix("test", "te"))
	p("HasSuffix:", ss.HasSuffix("test", "st"))
	p("Index: ", ss.Index("test", "e"))
	p("Join: ", ss.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat: ", ss.Repeat("sb", 5))
	p("Replace: ", ss.Replace("foo", "o", "0", -1))
	p("Replace: ", ss.Replace("foo", "o", "0", 1))
	p("Split: ", ss.Split("a-b-c-d-e", "-"))
	p("ToLower: ", ss.ToLower("TEST"))
	p("ToUpper: ", ss.ToUpper("test"))
}
