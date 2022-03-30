package pool

import "sync"

type Pool[T any] struct {
	pool sync.Pool
}

func (p *Pool[T]) Get() *T {
	return p.pool.Get().(*T)
}

func (p *Pool[T]) Put(elem *T) {
	p.pool.Put(elem)
}

func New[T any]() Pool[T] {
	return Pool[T]{
		pool: sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
}
