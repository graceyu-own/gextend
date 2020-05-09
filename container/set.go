package container

type Set interface {
	Collection
}

type AbstractSet struct {
	Set
	AbstractCollection
}
