package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
	//"strconv"
)

//struct - groups related data
type Task struct {
	Id    int
	Title string
	Done  bool
}
var tasks []Task
var taskIndex map[int]*Task
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
func AddTask(title string){
id := len(tasks)-1
task := Task{
	Id:id,
	Title:title,
	Done: false,
}
tasks = append(tasks,task)
taskIndex[id] = &tasks[len(tasks)-1]
}
func deleteTask(id int)bool{
	
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
task := Task{
	Title:"heelo",
	Id:32,
	Done:true,
}
// f
fmt.Println(task)
renameTask(&task,"main")
fmt.Println(task)

}
