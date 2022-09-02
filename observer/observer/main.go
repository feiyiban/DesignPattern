package main

import "fmt"

// ISubject
type ISubject interface {
	Register(obj IObject)
	Remove(obj IObject)
	Notify(msg string)
}

// Subject
type Subject struct {
	Obj []IObject
}

func (sub *Subject) Register(obj IObject) {
	sub.Obj = append(sub.Obj, obj)
}

func (sub *Subject) Remove(obj IObject) {
	for i, objValue := range sub.Obj {
		if objValue == obj {
			sub.Obj = append(sub.Obj[:i], sub.Obj[i+1:]...)
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, obj := range sub.Obj {
		obj.Update(msg)
	}
}

type IObject interface {
	Update(string)
}

type ObjOne struct {
}

func (obj *ObjOne) Update(msg string) {
	fmt.Printf("One say %v\n", msg)
}

type ObjTwo struct {
}

func (obj *ObjTwo) Update(msg string) {
	fmt.Printf("two say %v\n", msg)
}

type ObjThree struct {
}

func (obj *ObjThree) Update(msg string) {
	fmt.Printf("three say %v\n", msg)
}

func main() {
	sub := &Subject{}
	one := &ObjOne{}
	two := &ObjTwo{}
	three := &ObjThree{}

	sub.Register(one)
	sub.Register(two)
	sub.Register(three)

	sub.Notify("Hello")

	fmt.Println("**********************")

	sub.Remove(one)
	sub.Notify("world")
}
