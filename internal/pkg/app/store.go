package app

import (
	"fmt"

	v1 "github.com/arkadiusjonczek/go-rest-api-protobuf-protovalidate-demo.git/pkg/proto/demo/v1"
)

type Store[T any] interface {
	Add(entity *T) (int, error)
	Update(id int, entity *T) error
	Delete(id int) error
	Get(id int) (*T, error)
	GetAll() (map[int]*T, error)
}

var _ Store[*v1.Customer] = (*InMemoryStore[*v1.Customer])(nil)

type InMemoryStore[T any] struct {
	counter int
	entries map[int]*T
}

func NewInMemoryStore[T any]() *InMemoryStore[T] {
	return &InMemoryStore[T]{
		counter: 1,
		entries: make(map[int]*T),
	}
}

func (s *InMemoryStore[T]) Add(entry *T) (int, error) {
	id := s.counter

	s.entries[id] = entry

	s.counter = s.counter + 1

	return id, nil
}

func (s *InMemoryStore[T]) Update(id int, entry *T) error {
	s.entries[id] = entry

	return nil
}

func (s *InMemoryStore[T]) Delete(id int) error {
	s.entries[id] = nil

	return nil
}

func (s *InMemoryStore[T]) Get(id int) (*T, error) {
	entry, ok := s.entries[id]

	if !ok {
		return nil, fmt.Errorf("entry with id '%d' not found", id)
	}

	return entry, nil

}

func (s *InMemoryStore[T]) GetAll() (map[int]*T, error) {
	return s.entries, nil
}
