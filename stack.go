package stack

func New() *Stack {
	return &Stack()
}

type Stack struct{}

func (s Stack) IsEmpty() bool {
	return false
}
