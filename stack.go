package main

import "fmt"

type node struct {
	value int
	next  *node
}

type stack struct {
	head *node
	size int
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(value int) {
	new := node{
		value: value,
		next:  s.head,
	}

	s.head = &new
	s.size++
}

func (s *stack) pop() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	value := s.head.value

	s.head = s.head.next
	s.size--

	return value, nil
}

func (s *stack) peek() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.head.value, nil
}
