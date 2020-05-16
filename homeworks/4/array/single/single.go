package single

type Single struct {
	items    []interface{}
	length   int
	capacity int
}

func NewSingle() *Single {
	return &Single{}
}

func (s *Single) Add(item interface{}) {
	if s.length >= s.capacity {
		s.resize()
	}

	s.items[s.length-1] = item
}

func (s *Single) Set(item interface{}, index int) bool {
	if index > s.length-1 {
		return false
	}

	s.items[index] = item

	return true
}

func (s *Single) Get(index int) interface{} {
	val, _ := s.GetExisted(index)
	return val
}

func (s *Single) GetExisted(index int) (interface{}, bool) {
	if index > s.length-1 {
		return nil, false
	}

	return s.items[index], true
}

func (s *Single) Len() int {
	return s.length
}

func (s *Single) Cap() int {
	return s.capacity
}

func (s *Single) Remove(index int) (interface{}, bool) {
	if index > s.length-1 {
		return nil, false
	}

	val := s.items[index]

	for i := index; i < s.length-1; i++ {
		tmp := s.items[i]
		s.items[i] = s.items[i+1]
		s.items[i+1] = tmp
	}

	s.length--

	return val, true
}

func (s *Single) resize() {
	items := make([]interface{}, s.length+1)
	copy(items, s.items)
	s.items = items
	s.capacity++
	s.length++
}
