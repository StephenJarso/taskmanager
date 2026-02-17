package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Task struct {
	Id    int
	Title string
	Done  string
}

func showMenu() {
	fmt.Println("==Welcome to task manager==")
	fmt.Println("Press 1 to add task")
	fmt.Println("Press 2 to list tasks")
	fmt.Println("Press 3 to delete tasks")
	fmt.Println("Press 4 to search tasks")

	fmt.Print("Please select an option:")
}
func Addtasks(task *[]Task,reader *bufio.Reader){
	fmt.Print("Enter task number:")
	id,_:= reader.ReadString('\n')
	id = strings.TrimSpace(id)
	id1,_:= strconv.Atoi(id)
	fmt.Print("Enter task title:")
	title,_:= reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter task number:")
	done,_:= reader.ReadString('\n')
	done = strings.TrimSpace(done)
*task = append(*task,Task{
	Id:id1,
	Title:title,
	Done:done,
})
}
func main() {
	for {
	showMenu()
task := []Task{}
	reader := bufio.NewReader(os.Stdin)
	choice ,_:= reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	switch choice{
	case "1":
		Addtasks(&task, reader)

	default:
		fmt.Println("enter")
	}
}
}
