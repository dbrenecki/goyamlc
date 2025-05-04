package test

import "fmt"

// Example comment for animal
type Animal struct {
	Name string
	// Numb of legs
	Legs int
	// Species Species
}

var _ = somefunc

func somefunc() {
	fmt.Println("funcs are ignored")
}

// This is the root struct
// Second line for rootstruct
type Kingdom struct {
	// Some animal type.
	Animal Animal
	// Name represents the Kingdoms name.
	Name string
	// Plants []Plant
	Plant Plant
}

type Plant struct {
	Ipsom string
	Color []string
	// MiniPlant MiniPlant
}

// type Species struct {
// 	Something map[string]any
// 	Continent string
// 	IsExtinct bool
// }

// type MiniPlant struct {
// 	PlantType    string
// 	LeaveTypes   []LeaveType
// 	AnotherThing AnotherThing
// 	StrangeMap   map[string]Blah
// }

// type LeaveType struct {
// 	Color  string
// 	Color2 string
// }

// type AnotherThing struct {
// 	Foobar string
// }

// type Blah struct {
// 	Bar string
// }

// // This is another root struct
// type SecondRootStruct struct {
// 	Name string
// }
