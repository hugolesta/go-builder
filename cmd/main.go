package main

import (
	"github.com/hugolesta/go-builder/list"
	"github.com/hugolesta/go-builder/todo"
)

func main() {
	myTodos := factoryToDo("any")
	rendering := factoryList("plain")
	myTodos.SetList(rendering)
	myTodos.Add("Task1")
	myTodos.Add("Task2")
	myTodos.Add("Task3")
	myTodos.Add("Task4")
	myTodos.Add("Task5")
	myTodos.Print()
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
	return list.NewBullet('*')
}
