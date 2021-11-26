package maps

import (
	"sync"
)

type sorted struct {
	data map[interface{}]interface{}
	keys []interface{}

	mu sync.RWMutex
}

func NewSorted() *sorted {
	return &sorted{
		data: map[interface{}]interface{}{},
		keys: []interface{}{},
		mu:   sync.RWMutex{},
	}
}

func (m *sorted) Put(key, value interface{}, ex ...bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.put(key, value, ex...)
}

func (m *sorted) put(key, value interface{}, ex ...bool) {
	_, exist := m.data[key]
	if !exist {
		m.keys = append(m.keys, key)
	} else if len(ex) == 0 || !ex[0] {
		return
	}
	m.data[key] = value
}

func (m *sorted) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, exist := m.data[key]
	return v, exist
}

func (m *sorted) ContainsKey(key interface{}) bool {
	_, exist := m.Get(key)
	return exist
}

func (m *sorted) Remove(key interface{}) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.remove(key)
}

func (m *sorted) Removes(keys ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, key := range keys {
		m.remove(key)
	}
}

func (m *sorted) remove(key interface{}) (interface{}, bool) {
	v, exist := m.data[key]
	if exist {
		idx := -1
		for i := 0; i < len(m.keys); i++ {
			if m.keys[i] == key {
				idx = i
				break
			}
		}
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
		delete(m.data, key)
	}
	return v, exist
}

func (m *sorted) Replace(key, value interface{}, nx ...bool) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, exist := m.data[key]
	if exist {
		m.data[key] = value
	}
	if len(nx) > 0 && nx[0] {
		m.put(key, value, true)
	}
	return v, exist
}

func (m *sorted) Keys() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]interface{}, len(m.keys))
	copy(keys, m.keys)
	return keys
}

func (m *sorted) Values() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	vs := make([]interface{}, len(m.keys))
	for i := 0; i < len(m.keys); i++ {
		vs[i] = m.data[m.keys[i]]
	}
	return vs
}

func (m *sorted) ForEach(f func(key, value interface{}) bool) {
	nm := m.copy()
	for i := 0; i < len(nm.keys); i++ {
		k := nm.keys[i]
		if !f(k, nm.data[k]) {
			return
		}
	}
}

func (m *sorted) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

func (m *sorted) IsEmpty() bool {
	return m.Size() == 0
}

func (m *sorted) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = map[interface{}]interface{}{}
	m.keys = []interface{}{}
}

func (m *sorted) Copy() Map {
	return m.copy()
}

func (m *sorted) copy() *sorted {
	m.mu.RLock()
	defer m.mu.RUnlock()
	nm := NewSorted()
	for k, v := range m.data {
		nm.data[k] = v
	}
	nm.keys = make([]interface{}, len(m.keys))
	copy(nm.keys, m.keys)
	return nm
}
