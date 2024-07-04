package main

import (
	"fmt"
)

/*
*
泛型函数
*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k, v := range m {
		r = append(r, k)
		fmt.Println(k, v)
	}
	return r
}

/*
*
泛型类
*/
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	value T
	next  *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{v, nil}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{v, nil}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "1", 2: "2", 3: "3"}
	fmt.Println("keys = ", MapKeys(m))

	_ = MapKeys[int, string](m)

	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	fmt.Println(list.GetAll())
}
