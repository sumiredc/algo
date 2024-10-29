package data_structure

type Stack struct {
	data []int
	n    int
}

func NewStack() Stack {
	return Stack{n: -1}
}

func (s *Stack) Pop() *int {
	if s.Empty() {
		return nil
	}

	v := s.data[s.n]
	s.data = s.data[:s.n]
	s.n--

	return &v
}

func (s *Stack) Push(v int) {
	s.n++
	s.data = append(s.data, v)
}

func (s *Stack) Empty() bool {
	return s.n < 0
}

func (s *Stack) Len() int {
	return s.n + 1
}

func (s *Stack) Init() {
	s.n = -1
	s.data = nil
}
