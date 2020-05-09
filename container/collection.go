package container

import (
	"github.com/graceyu-own/gextend/lang"
	"github.com/graceyu-own/gextend/utils"
)

type Collection interface {
	Size() int
	IsEmpty() bool
	Contain(element lang.Object) bool
	Add(element lang.Object)
	Remove(element lang.Object) bool
	ToArray() []lang.Object
	Clear()

	utils.Iterable
	lang.Object
}

type AbstractCollection struct {
	Collection
}

func (_this *AbstractCollection) IsEmpty() bool {
	return _this.Size() == 0
}

func (_this *AbstractCollection) Contain(element lang.Object) bool {
	iterator := _this.Iterator()
	for iterator.HashNext() {
		if iterator.Next().Equals(element) {
			return true
		}
	}
	return false
}

func (_this *AbstractCollection) Remove(element lang.Object) bool {
	iterator := _this.Iterator()
	for iterator.HashNext() {
		if iterator.Next().Equals(element) {
			iterator.Remove()
			return true
		}
	}
	return false
}

func (_this *AbstractCollection) ToArray() []lang.Object {
	arr := make([]lang.Object, _this.Size())
	iterator := _this.Iterator()
	for iterator.HashNext() {
		arr = append(arr, iterator.Next())
	}
	return arr
}

func (_this *AbstractCollection) Clear() {
	iterator := _this.Iterator()
	for iterator.HashNext() {
		iterator.Next()
		iterator.Remove()
	}
}

func (_this *AbstractCollection) HashCode() int {
	code := 0
	iterator := _this.Iterator()
	for iterator.HashNext() {
		code += iterator.Next().HashCode()
	}
	return code
}
