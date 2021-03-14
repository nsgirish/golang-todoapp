package main 

import (
	"github.com/nsgirish/golang-todoapp/todotasks"
	"fmt"
)

func main() {
	fmt.Println("golang to do app")
	
	todotasks.AddTask("learn golang")
	todotasks.AddTask("learn dart")
	todotasks.PrintTasks()
}
