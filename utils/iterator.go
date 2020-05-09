package utils

import "github.com/graceyu-own/gextend/lang"

type Iterable interface {
	Iterator() Iterator
}

type Iterator interface {
	HashNext() bool
	Next() lang.Object
	Remove()
}
