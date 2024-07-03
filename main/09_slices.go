package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s), cap(s))

	s = make([]string, 3)
	fmt.Println("emp: ", s, s == nil, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: s = ", s)
	fmt.Println("get s[0] = ", s[0], ", len = ", len(s), cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append s = ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("slc: ", l)

	l = s[:5]
	fmt.Println("slc2: ", l)

	l = s[2:]
	fmt.Println("slc3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t, len(t), cap(t))

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range twoD[i] {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
