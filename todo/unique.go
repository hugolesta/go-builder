package todo

import "github.com/hugolesta/go-builder/list"

type Unique struct {
	rendering list.List
	todos     []string
}

func NewUnique() *Unique {
	return &Unique{}
}

func (u *Unique) SetList(l list.List) {
	u.rendering = l
}

func (u *Unique) Add(name string) {
	for _, v := range u.todos {
		if name == v {
			return
		}
	}
	u.todos = append(u.todos, name)
}

func (u Unique) Print() {
	u.rendering.Print(u.todos)
}
