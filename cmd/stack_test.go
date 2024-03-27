package main

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Add(3)
	stack.Add(5)
	stack.Add(7)

	if stack.Length != 3 {
		t.Error("Exptected Length to be 3")
	}
	if stack.Pop() != 7 {
		t.Error("Expected Pop() to return 7")
	}
	if stack.Pop() != 5 {
		t.Error("Expected Pop() to return 5")
	}
	if stack.Peek() != 3 {
		t.Error("Expected Peek() to return 3")
	}

	if stack.Pop() != 3 {
		t.Error("Expected Pop() to return 3")
	}
	if stack.Length != 0 {
		t.Error("Expected Length to be 0")
	}

	if stack.Pop() != nil {
		t.Error("Expected Pop() to return nil")
	}
	if stack.Peek() != nil {
		t.Error("Expected Peek() to return nil")
	}
}
