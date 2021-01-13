package todo

import (
	"container/list"

	_ "github.com/hugolesta/go-builder"
)

type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
