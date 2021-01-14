package todo

import (
	"github.com/hugolesta/go-builder/list"
)

type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
