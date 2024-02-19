package main

type stackElement[T any] struct {
	next *stackElement[T]
	val  T
}

type Stack[T any] struct {
	top *stackElement[T]
}

func (stack *Stack[T]) Push(v T) {
	el := &stackElement[T]{val: v}
	if stack.top == nil {
		stack.top = el
	} else {
		el.next = stack.top
		stack.top = el
	}
}

func (stack *Stack[T]) Pop(v T) {
	if stack.top != nil {
		stack.top = stack.top.next
	}
}

func (stack *Stack[T]) GetAll() []T {
	var elems []T

	for e := stack.top; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func stacks() {
}
