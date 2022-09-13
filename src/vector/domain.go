package vector

import "fmt"

type dynamicArray struct {
	elements []int
	count    int
}

func New() *dynamicArray {
	var v = &dynamicArray{}
	v.elements = make([]int, 16)
	v.count = 0

	return v
}

func (v *dynamicArray) Insert(element int, index int) {
	fmt.Println("Inserting ", element, " at ", index)
	if index >= v.count {
		panic("Index out of bounds")
	}

	if v.count == len(v.elements) {
		v.resize()
	}

	for i := v.count; i > index; i-- {
		v.elements[i] = v.elements[i-1]
	}

	v.elements[index] = element
	v.count++
}

func (v *dynamicArray) Prepend(element int) {
	fmt.Println("Prepending ", element)
	if v.count == len(v.elements) {
		v.resize()
	}

	for i := v.count; i > 0; i-- {
		v.elements[i] = v.elements[i-1]
	}

	v.elements[0] = element
	v.count++
}

func (v *dynamicArray) Pop() int {
	fmt.Println("Popping ", v.elements[v.count-1])
	if v.count == 0 {
		panic("Empty array")
	}

	var element = v.elements[v.count-1]
	v.elements[v.count-1] = 0
	v.count--
	return element
}

func (v *dynamicArray) Delete(index int) {
	fmt.Println("Deleting element at ", index)
	if index >= v.count {
		panic("Index out of bounds")
	}

	for i := index; i < v.count-1; i++ {
		v.elements[i] = v.elements[i+1]
	}

	v.count--
}

func (v *dynamicArray) Remove(element int) {
	fmt.Println("Removing ", element)
	for i := 0; i < v.count; i++ {
		if v.elements[i] == element {
			v.Delete(i)
			i--
		}
	}
}

func (v *dynamicArray) Find(element int) int {
	fmt.Println("Finding ", element)
	for i := 0; i < v.count; i++ {
		if v.elements[i] == element {
			return i
		}
	}

	return -1
}

func (v *dynamicArray) Print() {
	fmt.Printf("%v+\n", v)
}

func (v *dynamicArray) Capacity() int {
	return len(v.elements)
}

func (v *dynamicArray) IsEmpty() bool {
	return v.count == 0
}

func (v *dynamicArray) At(index int) int {
	fmt.Println("Element at  ", index, " is ", v.elements[index])
	if index >= v.count {
		panic("Index out of bounds")
	}

	return v.elements[index]
}

func (v *dynamicArray) Push(element int) {
	fmt.Println("Pushing ", element)
	if v.count == len(v.elements) {
		v.resize()
	}

	v.elements[v.count] = element
	v.count++
}

//private

func (v *dynamicArray) resize() {
	fmt.Println("resizing to ", len(v.elements)*2)
	var newElements = make([]int, len(v.elements)*2)
	copy(newElements, v.elements)
	v.elements = newElements
}
