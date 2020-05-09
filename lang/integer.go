package lang

import "strconv"

type Integer struct {
	AbstractObject
	int
}

func NewInteger(value int) Integer {
	return Integer{AbstractObject{}, value}
}

func (_this Integer) HashCode() int {
	code := _this.int
	return code
}

func (_this Integer) ToString() String {
	return NewString(strconv.Itoa(_this.int))
}

func (_this Integer) Value() interface{} {
	return _this.int
}
