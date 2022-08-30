package main

import "fmt"

// create a new type name Stack, that is a struct
// data is a slice of interfaces --> elements of Stack
// Khi khoi tao interface nhu vay se giup cho cac phan tu trong
// 1 slice nam ben tronf interface nen khong can chung 1 type
type Stack struct {
	data []interface{}
	top  int
}

// implementation of Push
func (s *Stack) Push(element interface{}){
	s.top++
	s.data = append(s.data, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		s.top--
		last := s.data[s.top]
		s.data = s.data[:s.top]

		return last
	}

	return nil
}

func main() {
	stack := Stack{}
	stack.Push("Do Tieu Thien")
	stack.Push(1000)
	stack.Push("Do Tieu Thien")
	stack.Push("Do Tieu Khoi")
	stack.Pop()

	fmt.Println(stack)
}
