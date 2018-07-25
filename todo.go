package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type List struct {
	Task []todo
}
type todo struct {
	Content string
	Done    bool
	Num     int
}

func (l *List) list() string {
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
func (l *List) delete(id int) error {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task = append(l.Task[:i], l.Task[i+1:]...)
			return nil
		}

	}
	return fmt.Errorf("ID does not exist")
}

func (l *List) done(id int) error {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task[i].Done = true
			return nil

		}
	}
	return fmt.Errorf("No task found with ID[%v],use a valid ID ", id)
}
func (l *List) add(s string) {
	max := 0
	for _, t := range l.Task {
		if t.Num > max {
			max = t.Num
		}
	}
	l.Task = append(l.Task, todo{Content: s, Num: max + 1})
}
func main() {
	l := List{}
	l.load()

	switch os.Args[1] {

	default:
		log.Fatalln("Invalid command, todo H for for a list of commands")
	case "Add", "add", "A", "a":
		l.add(strings.Join(os.Args[2:], " "))

	case "Remove", "remove", "R", "r":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid ID [%v],Input an ID from the list to remove", os.Args[2])
		}
		err = l.delete(id)
		if err != nil {
			fmt.Println("ID does not exist")
		}

	case "Done", "done", "D", "d":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("ID [%v] invalid,use valid ID ", os.Args[2])
		}
		err = l.done(id)
		if err != nil {
			log.Fatalln(err)
		}
	case "List", "list", "L", "l":
		fmt.Print(l.list())

	case "Help", "help", "H", "h":
		fmt.Println("todo L: prints the todo-list")
		fmt.Println("todo D [ID]: marks selected ID as done")
		fmt.Println("todo R [ID]: removes selected ID")
		fmt.Println("todo A [UserInput]: adds task to the list")
	}
	l.save()
}

func (l *List) save() {
	data, err := json.Marshal(l)
	if err != nil {
		log.Fatalln("Failed to save")
	}
	err = ioutil.WriteFile("todo.json", data, 0644)
	if err != nil {
		log.Fatalln("Failed to write.")
	}
}
func (l *List) load() {
	dat, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Fatalln("Failed to load.")

	}
	err = json.Unmarshal(dat, l)
	if err != nil {
		log.Fatalln("Failed to unmartial.")
	}

}
