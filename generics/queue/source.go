package queue

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Sth=string,int,float64"

type Sth generic.Type

type SthQueue struct {
	items []Sth
}

func MewSthQueue() *SthQueue {
	return &SthQueue{
		items: make([]Sth, 0),
	}
}

func (q *SthQueue) Push(item Sth) {
	q.items = append(q.items, item)
}

func (q *SthQueue) Pop() Sth {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
