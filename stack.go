package stack

import (
    "fmt"
    "strings"
)

type PointerStackNode struct {
    value string
    next *PointerStackNode
}

type PointerStack struct {
    head *PointerStackNode
}

func NewPointerStack(initValue []string) *PointerStack {
    result := new(PointerStack)
    for _, value := range initValue {
        result.Push(value)
    }
    return result
}

func (s *PointerStack) Push(v string) {
    s.head = &PointerStackNode{value: v, next: s.head}
}

func (s *PointerStack) Pop() (string, error) {
    if s.head == nil {
        return "", emptyOperationError("pop")
    }
    result := s.head.value
    s.head = s.head.next
    return result, nil
}

func (s *PointerStack) Peek() (string, error) {
    if s.head == nil {
        return "", emptyOperationError("peek")
    }
    return s.head.value, nil
}

func (s *PointerStack) Print() (result string) {
    var values []string
    node := s.head
    for node != nil {
        values = append(values, node.value)
        node = node.next
    }
    return "(top)[" + strings.Join(values, ",") + "](bottom)"
}

func (s *PointerStack) Reverse() {
    var prevNode *PointerStackNode
    node := s.head

    for node != nil {
        nextNode := node.next
        s.head = node
        s.head.next = prevNode
        prevNode = s.head
        node = nextNode
    }
}

func (s *PointerStack) IsEmpty() bool {
    return s.head == nil
}

func emptyOperationError(operation string) error {
    return fmt.Errorf("unable to %s: Stack is empty", operation)
}