package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
)
var tasks []Task
var taskIndex map[int]*Task

type TaskManager struct{
	tasks []Task
	taskIndex map[int]*Task
}
//struct - groups related data
type Task struct {
	Id    int
	Title string
	Done  bool
}

func showMenu() {
	fmt.Println("==Welcome to task manager==")
	fmt.Println("Press 1 to add task")
	fmt.Println("Press 2 to list tasks")
	fmt.Println("Press 3 to delete tasks")
	fmt.Println("Press 4 to search tasks")

	fmt.Print("Please select an option:")

}
//storage
func initStorage(){
	tasks = make([]Task,0)
	taskIndex = make(map[int]*Task)
	}
func AddTask(tasks *[]Task, reader *bufio.Reader){
id := len(*tasks)-1
fmt.Print("Enter task title:")
title,_:= reader.ReadString('\n')
title = strings.TrimSpace(title)
*tasks = append(*tasks,Task{
	Title:title,
	Id:id,
	Done:false,
})
}
func deleteTask(id int)bool{
if taskIndex[id] == nil {
	return false
}	
index :=-1
for i := range tasks{
	if tasks[i].Id == id{
		index = i
		break
	}
}
if index == -1{
	return false
}
//Remove from slice
 tasks = append(tasks[:index], tasks[index+1:]...)

    // Remove from map
    delete(taskIndex, id)

    // Rebuild map pointers
    for i := range tasks {
        taskIndex[tasks[i].Id] = &tasks[i]
    }

    return true
}
// func findTask(id int)*Task{
// 	for i := range tasks{
// 		if tasks[id].Id == id{
// 			return &tasks[i]
// 		}
// 	}
// 	return nil
// }
func renameTask(t *Task,title string){
	t.Title = title
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	choice,_ :=reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
task := Task{
	Title:"heelo",
	Id:32,
	Done:true,
}

showMenu()
switch choice{
case "1":
	AddTask(&tasks,reader)
case "2":
	fmt.Println("press 2 to exit")
	return
}

renameTask(&task,"main")
fmt.Println(task)


}
