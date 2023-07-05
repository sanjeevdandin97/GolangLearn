package goReturnType

import "fmt"

var GlobalString string   // output: ""
var GlobalInteger int     // output: 0
var GlobalFloat32 float32 // output: 0
var GlobalFloat64 float64 // output: 0
var GlobalBoolean bool    // output: false

func goReturnString() string {
	localString := GlobalString
	return localString
}

func goReturnInteger() int {
	localString := GlobalInteger
	return localString
}

func goReturnFloat32() float32 {
	localFloat32 := GlobalFloat32
	return localFloat32
}

func goReturnFloat64() float64 {
	localFloat64 := GlobalFloat64
	return localFloat64
}

func goReturnBoolean() bool {
	localBoolean := GlobalBoolean
	return localBoolean
}

func GoReturnTypePrint() {
	fmt.Println(goReturnString())  // STRING
	fmt.Println(goReturnInteger()) // INTEGER
	fmt.Println(goReturnFloat32()) // FLOAT 32 bit
	fmt.Println(goReturnFloat64()) // FLOAT 64 bit
	fmt.Println(goReturnBoolean()) // BOOLEAN
}
