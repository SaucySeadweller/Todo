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
	task []todo
}
type todo struct {
	content string
	done    bool
	num     int
}

func (l *List) list() string {
	s := ""
	for _, t := range l.task {
		if t.done {
			log.Println("Done")
			s += fmt.Sprintf("%v. %v [Done]\n", t.num, t.content)
		} else {
			s += fmt.Sprintf("%v. %v\n", t.num, t.content)
		}
	}
	return s
}
func (l *List) delete(id int) {
	for i, t := range l.task {
		if t.num == id {
			l.task = append(l.task[:i], l.task[i+1])
			return
		}
	}
}

func (l *List) done(id int) {

	for i, t := range l.task {
		if t.num == id {
			l.task[i].done = true
			break
		}
	}
}
func (l *List) add(s string) {
	max := 0
	for _, t := range l.task {
		if t.num > max {
			max = t.num
		}
	}
	l.task = append(l.task, todo{content: s, num: max + 1})
}
func main() {
	l := List{}
	l.load()
	switch os.Args[1] {
	case "add":
		l.add(strings.Join(os.Args[2:], " "))
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		l.delete(id)
	case "done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		l.done(id)
	case "list":
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
		log.Println(err)
		return
	}
	err = json.Unmarshal(dat, l)
	if err != nil {
		panic(err)
	}

}
