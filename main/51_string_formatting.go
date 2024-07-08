package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{
		x: 1,
		y: 2,
	}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool %t\n", true)
	fmt.Printf("int %d\n", 123)
	fmt.Printf("bin %b\n", 14)
	fmt.Printf("char %c\n", 97)
	fmt.Printf("hex 0x%x\n", 456)

	fmt.Printf("float1 %.3f\n", 123.456)
	fmt.Printf("float2 %e\n", 123.456)
	fmt.Printf("float3 %E\n", 123.456)

	fmt.Printf("string1 %s\n", "\"hello\"")
	fmt.Printf("string2 %q\n", "\"hello\"")
	fmt.Printf("string3 %x\n", "\"hello\"") // hex string

	fmt.Printf("pointer %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1234567.2, 3.456789)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1234567.2, 3.456789)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
