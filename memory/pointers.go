package main

import (
	"fmt"
)

type IntContainer struct {
	x int
}

func main() {
	pointerDereferncing()
	stackAllocatedPointer()
}

func pointerDereferncing() {
	fmt.Println("**** Start pointerDereferncing ****")
	container := IntContainer{x: 3}
	pointer := &container
	fmt.Println("Direct access:", container.x)
	fmt.Println("Implicit dereference:", pointer.x)
	fmt.Println("Explicit dereference", (*pointer).x)
	fmt.Println("**** End pointerDereferncing ****\n\n")
}

func stackAllocatedPointer() {
	fmt.Println("**** Start stackAllocatedPointer ****")
	pointer_to_stack := stackPointer()
	fmt.Printf("Address on heap:    %p\n", pointer_to_stack)
	fmt.Println(pointer_to_stack.x)
	fmt.Println("**** End stackAllocatedPointer ****\n\n")
}

func stackPointer() *IntContainer {
	// IntContainer{5} makes a new IntContainer
	// with x=5 that theoretically exists on the stack
	container := IntContainer{x: 5}
	fmt.Printf("Address \"on stack\": %p\n", &container)
	// we can return a pointer to this stack variable
	// the compiler will realize that we are returning
	// a pointer to a struct on the stack, it will thus
	// allocate the struct on the heap and return a
	// safe pointer
	return &container
}