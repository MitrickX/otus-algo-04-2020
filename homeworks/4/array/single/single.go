package single

type Single struct {
	items    []interface{}
	length   int
	capacity int
}

func NewSingle() *Single {
	return &Single{}
}

func (s *Single) Append(item interface{}) {
	if s.length >= s.capacity {
		s.resize()
	}

	s.items[s.length-1] = item
}

func (s *Single) Add(item interface{}, index int) bool {
	if index > s.length {
		return false
	}

	if index == s.length {
		s.Append(item)
		return true
	}

	if s.length >= s.capacity {
		items := s.allocate()
		copy(items[0:index], s.items[0:index])
		items[index] = item
		copy(items[index+1:s.length+1], s.items[index:s.length])
		s.items = items
		s.length++
		s.capacity++

		return true
	}

	copy(s.items[index+1:s.length+1], s.items[index:s.length])
	s.items[index] = item
	s.length++

	return true
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

func (s *Single) allocate() []interface{} {
	return make([]interface{}, s.length+1)
}

func (s *Single) resize() {
	items := s.allocate()
	copy(items, s.items)
	s.items = items
	s.capacity++
	s.length++
}
