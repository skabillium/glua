package main

type StackNode struct {
	Value any
	Prev  *StackNode
}

type Stack struct {
	Length int
	top    *StackNode
}

func NewStack() *Stack {
	return &Stack{Length: 0, top: nil}
}

func (s *Stack) Add(item any) {
	s.Length++
	node := &StackNode{Value: item}
	if s.Length == 1 {
		s.top = node
		return
	}

	node.Prev = s.top
	s.top = node
}

func (s *Stack) Pop() any {
	if s.Length == 0 {
		return nil
	}

	s.Length--
	if s.Length == 0 {
		value := s.top.Value
		s.top = nil
		return value
	}

	value := s.top.Value
	s.top = s.top.Prev
	return value
}

func (s *Stack) Peek() any {
	if s.Length == 0 {
		return nil
	}

	return s.top.Value
}
