package container

import (
	"github.com/graceyu-own/gextend/lang"
)

type List interface {
	Collection
}

type AbstractList struct {
	AbstractCollection
	loadFactor float64
	capacity   int
	number     int
}

type ArrayList struct {
	AbstractList
	elements []lang.Object
}

func NewArrayList() ArrayList {
	a := ArrayList{AbstractList{
		AbstractCollection{new(ArrayList)}, 0.75, 16, 0},
		make([]lang.Object, 16),
	}
	return a
}
