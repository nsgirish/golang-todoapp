package todotasks

import (
"fmt"
)

var tasks = make(map[string]bool)

func AddTask(title string) {
	tasks[title] = false
}

func PrintTasks() {
	for k,v := range tasks {
		fmt.Println("title:",k,"completed:",v)
	}
}
