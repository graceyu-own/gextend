package lang

type String struct {
	AbstractObject
	string
}

func NewString(str string) String {
	return String{AbstractObject{}, str}
}

func (_this String) HashCode() int {

	code := 0

	for _, b := range []byte(_this.string)[:] {
		code += int(b)
	}

	return code
}

func (_this String) ToString() String {
	return _this
}

func (_this String) Value() interface{} {
	return _this.string
}
