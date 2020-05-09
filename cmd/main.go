package main

import (
	"github.com/graceyu-own/gextend/container"
	"github.com/graceyu-own/gextend/lang"
)

func main() {

	list := container.NewArrayList()
	list.Add(lang.NewInteger(5))
	list.Remove(lang.NewInteger(5))

}
