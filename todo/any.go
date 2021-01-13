package todo

import "github.com/hugolesta/go-builder"

type Any stuct {
	rendering list.List
	todos []string
}

fun NewAny() *Any {
	return &Any{}
}

func (a *Any) SetList(l list.List){
	a.rendering = l
}

func (a *Any) Add(name string) {
	a.todos = append(a.todos, name)
}

func (a *Any) Print() {
	a.rendering.Print(a.todos)
}