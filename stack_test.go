package main

import (
	"testing"
)

func assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}

func testStack_push(t *testing.T, input []int, expectedSize int) {
	s := newStack()

	for _, value := range input {
		s.push(value)
	}

	assert(t, s.size, expectedSize)

	l := len(input)
	currentNode := s.head

	for i := 0; i < l; i++ {
		assert(t, currentNode.value, input[l-i-1])
		currentNode = currentNode.next
	}

	assert(t, currentNode, (*node)(nil))
}

func TestStack_init(t *testing.T) {
	testStack_push(t, []int{}, 0)
}

func TestStack_push_1(t *testing.T) {
	input := []int{5}
	testStack_push(t, input, 1)
}

func TestStack_push_3(t *testing.T) {
	input := []int{5, 4, 3}
	testStack_push(t, input, 3)
}

func testStack_pop(
	t *testing.T,
	input []int,
	popsSize int,
	expectedValues []int,
	expectedErrs []bool) {

	s := &stack{}

	for _, value := range input {
		s.push(value)
	}

	for i := 0; i < popsSize; i++ {
		actualValue, actualErr := s.pop()

		assert(t, actualValue, expectedValues[i])
		assert(t, actualErr != nil, expectedErrs[i])
	}
}

func TestStack_pop_empty(t *testing.T) {
	input := []int{}
	popsSize := 1
	expectedValues := []int{0}
	expectedErrs := []bool{true}

	testStack_pop(t, input, popsSize, expectedValues, expectedErrs)
}

func TestStack_pop_once(t *testing.T) {
	input := []int{7}
	popsSize := 1
	expectedValues := []int{7}
	expectedErrs := []bool{false}

	testStack_pop(t, input, popsSize, expectedValues, expectedErrs)
}

func TestStack_pop_twice(t *testing.T) {
	input := []int{7}
	popsSize := 2
	expectedValues := []int{7, 0}
	expectedErrs := []bool{false, true}

	testStack_pop(t, input, popsSize, expectedValues, expectedErrs)
}
