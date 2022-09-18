package dataStructures

type stack []string

type Stack interface {
	Push(element string)
	Pop() (string, error)
}

func NewStack() stack {
	return stack{}
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(element string) {
	*s = append(*s, element)
}

func (s *stack) Pop() (string, error) {
	if s.isEmpty() {
		return "", ErrStackEmpty
	}
	index := len(*s)
	element := (*s)[index-1]
	*s = (*s)[:index-1]
	return element, nil
}
