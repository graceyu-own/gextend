package lang

type Object interface {
	HashCode() int
	ToString() String
	Equals(obj Object) bool
}

type AbstractObject struct {
	Object
}

func (_this AbstractObject) Equals(obj Object) bool {
	return _this == obj
}
