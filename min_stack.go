package main

import (
	"fmt"
	"math"
)

type minStack struct {
	stack
	currentMin int
	allMins    *stack
}

func newMinStack() *minStack {
	return &minStack{
		currentMin: math.MaxInt32,
		allMins:    newStack(),
	}
}

func (ms *minStack) push(value int) {
	ms.stack.push(value)

	if value <= ms.currentMin {
		ms.allMins.push(value)
		ms.currentMin = value
	}
}

func (ms *minStack) pop() (int, error) {
	value, err := ms.stack.pop()

	if err == nil && value == ms.currentMin {
		ms.allMins.pop()

		nextMin, err := ms.allMins.peek()

		if err != nil {
			nextMin = math.MaxInt32
		}

		ms.currentMin = nextMin
	}

	return value, err
}

func (ms *minStack) min() (int, error) {
	if ms.head == nil {
		return math.MaxInt32, fmt.Errorf("stack is empty")
	}

	return ms.currentMin, nil
}
