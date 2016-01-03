package api

import "fmt"

func Store(todo Todo) Todo  {
	fmt.Printf("%s\n", todo.Name)
	return todo
}