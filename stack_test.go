package stack

import (
    "testing"
)

func TestNewPointerStack(t *testing.T) {
    stack := NewPointerStack([]string{"1", "2"})
    expectedValues := []string{"2", "1"}
    for _, expected := range expectedValues {
        if stack.head.value != expected {
            t.Fatalf("expected %s, got %s", expected, stack.head.value)
        }
        stack.head = stack.head.next
    }
}

func TestPointerStack_Pop(t *testing.T) {
    stack := NewPointerStack([]string{"1", "2"})
    expectedValues := []string{"2", "1"}
    for _, expected := range expectedValues {
        if result, _ := stack.Pop(); result != expected {
            t.Fatalf("expected %s, got %s", expected, result)
        }
    }
}

func TestPointerStack_Pop_Empty(t *testing.T) {
    stack := new(PointerStack)

    if _, err := stack.Pop(); err == nil {
        t.Fatal("expected error, got nil")
    }
}

func TestPointerStack_Print(t *testing.T) {
    stack := NewPointerStack([]string{"1", "2"})
    expected := "(top)[2,1](bottom)"

    if result := stack.Print(); result != expected {
        t.Fatalf("expected %s, got %s", expected, result)
    }
}

func TestPointerStack_Reverse(t *testing.T) {
    stack := NewPointerStack([]string{"1", "2"})
    expected := "(top)[1,2](bottom)"

    stack.Reverse()

    if result := stack.Print(); result != expected {
        t.Fatalf("expected %s, got %s", expected, result)
    }
}