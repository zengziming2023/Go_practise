package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPoint(pval *int) {
	*pval = 0
}

func main() {
	i := 1
	fmt.Println("init i = ", i)
	fmt.Println("init i address = ", &i)
	zeroVal(i)
	fmt.Println("zeroVal i = ", i)

	zeroPoint(&i)
	fmt.Println("zeroPoint i = ", i)

	fmt.Println("pointer i = ", &i)
}
