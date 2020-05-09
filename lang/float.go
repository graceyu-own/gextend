package lang

import "strconv"

type Float32 struct {
	AbstractObject
	float32
}

func (_this Float32) HashCode() int {
	hashCode := 0
	hashCode = int(_this.float32)
	return hashCode
}

func (_this Float32) ToString() String {
	return NewString(strconv.FormatFloat(float64(_this.float32), 'f', 6, 32))
}

type Float64 struct {
	AbstractObject
	float64
}

func (_this Float64) HashCode() int {
	hashCode := 0
	hashCode = int(_this.float64)
	return hashCode
}

func (_this Float64) ToString() String {
	return NewString(strconv.FormatFloat(_this.float64, 'f', 6, 64))
}
