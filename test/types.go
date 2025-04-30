package test

import "fmt"

type Animal struct {
	Name    string
	Legs    int
	Species Species
}

var _ = somefunc

func somefunc() {
	fmt.Println("funcs are ignored")
}

type Kingdom struct {
	// Some animal
	Animal Animal
	Name   string
	Plants []Plant
}

type Plant struct {
	Ipsom     string
	Color     []string
	MiniPlant MiniPlant
}

type Species struct {
	Something map[string]any
	Continent string
	IsExtinct bool
}

type MiniPlant struct {
	PlantType    string
	LeaveTypes   []LeaveTypes
	AnotherThing AnotherThing
	StrangeMap   map[string]Blah
}

type LeaveTypes struct {
	Color string
}

type AnotherThing struct {
	Foobar string
}

type Blah struct {
	Bar string
}
