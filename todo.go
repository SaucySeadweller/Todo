package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
func (l *List) delete(id int) {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task = append(l.Task[:i], l.Task[i+1:]...)
			return
		}
	}
}

func (l *List) done(id int) {
	for i, t := range l.Task {
		if t.Num == id {
			l.Task[i].Done = true
			break
		}
	}
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
	case "add", "a":
		l.add(strings.Join(os.Args[2:], " "))
	case "remove", "r":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		l.delete(id)
	case "done", "d":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		l.done(id)
	case "list", "l":
		fmt.Print(l.list())
	}
	l.save()
}
func (l *List) save() {
	data, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("todo.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
func (l *List) load() {
	dat, err := ioutil.ReadFile("todo.json")
	if err != nil {

		return
	}
	err = json.Unmarshal(dat, l)
	if err != nil {
		panic(err)
	}

}
