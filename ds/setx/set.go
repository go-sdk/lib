package setx

import (
	"sync"
)

type Set interface {
	Add(value interface{}) bool
	Adds(values ...interface{}) int
	Contains(values ...interface{}) bool
	Remove(value interface{}) bool
	Removes(values ...interface{}) int
	Values() []interface{}
	ForEach(f func(value interface{}) bool)

	Size() int
	IsEmpty() bool
	Clear()
	Copy() Set
}

type set struct {
	data     []interface{}
	valueMap map[interface{}]struct{}

	mu sync.RWMutex
}

func NewSet() *set {
	return &set{
		data:     []interface{}{},
		valueMap: map[interface{}]struct{}{},
		mu:       sync.RWMutex{},
	}
}

func (s *set) Add(value interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.add(value)
}

func (s *set) Adds(values ...interface{}) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for i := 0; i < len(values); i++ {
		if s.add(values[i]) {
			count++
		}
	}
	return count
}

func (s *set) add(value interface{}) bool {
	_, exist := s.valueMap[value]
	if exist {
		return false
	}
	s.data = append(s.data, value)
	s.valueMap[value] = struct{}{}
	return true
}

func (s *set) Contains(values ...interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i := 0; i < len(values); i++ {
		_, exist := s.valueMap[values[i]]
		if exist {
			return true
		}
	}
	return false
}

func (s *set) Remove(value interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.remove(value)
}

func (s *set) Removes(values ...interface{}) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for i := 0; i < len(values); i++ {
		if s.remove(values[i]) {
			count++
		}
	}
	return count
}

func (s *set) remove(value interface{}) bool {
	_, exist := s.valueMap[value]
	if !exist {
		return false
	}
	idx := -1
	for i := 0; i < len(s.data); i++ {
		if s.data[i] == value {
			idx = i
			break
		}
	}
	s.data = append(s.data[:idx], s.data[idx+1:]...)
	delete(s.valueMap, value)
	return true
}

func (s *set) Values() []interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	vs := make([]interface{}, len(s.data))
	copy(vs, s.data)
	return vs
}

func (s *set) ForEach(f func(value interface{}) bool) {
	ns := s.copy()
	for i := 0; i < len(ns.data); i++ {
		if !f(ns.data[i]) {
			return
		}
	}
}

func (s *set) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *set) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = []interface{}{}
	s.valueMap = map[interface{}]struct{}{}
}

func (s *set) Copy() Set {
	return s.copy()
}

func (s *set) copy() *set {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ns := NewSet()
	ns.data = make([]interface{}, len(s.data))
	copy(ns.data, s.data)
	for v := range s.valueMap {
		ns.valueMap[v] = struct{}{}
	}
	return ns
}
