package main

import (
	"app/src/vector"
)

func main() {
	array := vector.New()
	for i := 0; i < 100; i++ {
		array.Push(i)
	}

	array.Pop()
	array.Push(111)
	array.Print()
}
