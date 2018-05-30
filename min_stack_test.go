package main

import (
	"math"
	"testing"
)

func testMinStack(
	t *testing.T,
	input []int,
	popsSize int,
	expectedValue int,
	expectedErr bool) {

	ms := newMinStack()

	for _, value := range input {
		ms.push(value)
	}

	for i := 0; i < popsSize; i++ {
		ms.pop()
	}

	actualValue, actualErr := ms.min()

	assert(t, expectedValue, actualValue)
	assert(t, expectedErr, actualErr != nil)
}

func TestMinStack_min_empty(t *testing.T) {
	input := []int{}
	popsSize := 0
	expectedValue := math.MaxInt32
	expectedErr := true

	testMinStack(t, input, popsSize, expectedValue, expectedErr)
}

func TestMinStack_min_single(t *testing.T) {
	input := []int{5}
	popsSize := 0
	expectedValue := 5
	expectedErr := false

	testMinStack(t, input, popsSize, expectedValue, expectedErr)
}

func TestMinStack_min_triple(t *testing.T) {
	input := []int{5, 3, 2}
	popsSize := 1
	expectedValue := 3
	expectedErr := false

	testMinStack(t, input, popsSize, expectedValue, expectedErr)
}

func TestMinStack_min_identical(t *testing.T) {
	input := []int{5, 3, 2, 3, 2}
	popsSize := 1
	expectedValue := 2
	expectedErr := false

	testMinStack(t, input, popsSize, expectedValue, expectedErr)
}

func TestMinStack_min_pop_empty(t *testing.T) {
	input := []int{5, 3}
	popsSize := 2
	expectedValue := math.MaxInt32
	expectedErr := true

	testMinStack(t, input, popsSize, expectedValue, expectedErr)
}
