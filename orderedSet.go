package lift

import "sort"

type FloorRequest struct {
	Source      int
	Destination int
}

type OrderedSet struct {
	data  map[FloorRequest]struct{}
	items []FloorRequest
}

// Creates a OrderedSet
func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		data:  make(map[FloorRequest]struct{}),
		items: []FloorRequest{},
	}
}

func (s *OrderedSet) Add(req FloorRequest) {
	if _, exists := s.data[req]; !exists {
		s.data[req] = struct{}{}
		index := sort.Search(len(s.items), func(i int) bool { return s.items[i].Source >= req.Source })
		s.items = append(s.items[:index], append([]FloorRequest{req}, s.items[index:]...)...)

	}
}

func (s *OrderedSet) Remove(req FloorRequest) {
	if _, exists := s.data[req]; !exists {
		s.data[req] = struct{}{}
		index := sort.Search(len(s.items), func(i int) bool { return s.items[i].Source >= req.Source })
		if index < len(s.items) && s.items[index].Source == req.Source {
			s.items = append(s.items[:index], s.items[index+1:]...)
		}
	}
}

func (s *OrderedSet) Contains(item FloorRequest) bool {
	_, exists := s.data[item]
	return exists
}

func (s *OrderedSet) Items() []FloorRequest {
	return s.items
}

func (s *OrderedSet) Size() int {
	return len(s.data)
}
