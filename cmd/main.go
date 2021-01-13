package main

import (
	"github.com/hugolesta/go-builder/list"
	"github.com/hugolesta/go-builder/todo"
)

func main() {
	myTodos := factoryToDo("any")
	rendering := factoryList("plain")
	myTodos.SetList(rendering)
	muTodos.Add("Task1")
	muTodos.Add("Task2")
	muTodos.Add("Task3")
	muTodos.Add("Task4")
	muTodos.Add("Task5")
}

func factoryToDo(s string) todo.ToDo {
	if s == "unique" {
		return todo.NewUnique()
	}
	return todo.NewAny()
}

func factoryList(s string) list.List {
	if s == "plain" {
		return list.NewPlain()
	}
	return list.NewBullet("*")
}
