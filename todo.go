package Todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type TodoList struct {
	Task []Todo
}
type Todo struct {
	Content string
	Done    bool
	Num     int
}

func (l *TodoList) List() string {
	s := ""
	for _, t := range l.Task {
		if t.Done {
			s += fmt.Sprintf("%v. %v [Done]\n", t.Num, t.Content)
		} else {
			s += fmt.Sprintf("%v. %v\n", t.Num, t.Content)
		}
	}
	return s
}
func (l *TodoList) Delete(id int) error {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task = append(l.Task[:i], l.Task[i+1:]...)
			return nil
		}

	}
	return fmt.Errorf("ID does not exist")
}

func (l *TodoList) Done(id int) error {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task[i].Done = true
			return nil

		}
	}
	return fmt.Errorf("No task found with ID[%v],use a valid ID ", id)
}
func (l *TodoList) Add(s string) {
	max := 0
	for _, t := range l.Task {
		if t.Num > max {
			max = t.Num
		}
	}
	l.Task = append(l.Task, Todo{Content: s, Num: max + 1})
}
func main() {
	l := TodoList{}
	l.Load()
	if len(os.Args) == 2 {
		fmt.Println("")
	}

	switch os.Args[1] {

	default:
		log.Fatalln("Invalid command, todo H for for a list of commands")
	case "todo":
		log.Fatalf("nope")
	case "Add", "add", "A", "a":
		l.Add(strings.Join(os.Args[2:], " "))

	case "Remove", "remove", "R", "r":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid ID [%v],Input an ID from the list to remove", os.Args[2])
		}
		err = l.Delete(id)
		if err != nil {
			fmt.Println("ID does not exist")
		}

	case "Done", "done", "D", "d":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("ID [%v] invalid,use valid ID ", os.Args[2])
		}
		err = l.Done(id)
		if err != nil {
			log.Fatalln(err)
		}
	case "List", "list", "L", "l":
		fmt.Print(l.List())

	case "Help", "help", "H", "h,":
		fmt.Println("todo L: prints the todo-list")
		fmt.Println("todo D [ID]: marks selected ID as done")
		fmt.Println("todo R [ID]: removes selected ID")
		fmt.Println("todo A [UserInput]: adds task to the list")
	}
	l.Save()
}

func (l *TodoList) Save() {
	data, err := json.Marshal(l)
	if err != nil {
		log.Fatalln("Failed to save")
	}
	err = ioutil.WriteFile("todo.json", data, 0644)
	if err != nil {
		log.Fatalln("Failed to write.")
	}
}
func (l *TodoList) Load() {
	dat, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Fatalln("Failed to load.")

	}
	err = json.Unmarshal(dat, l)
	if err != nil {
		log.Fatalln("Failed to unmartial.")
	}

}
