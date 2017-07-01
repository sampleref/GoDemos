package main

import (
	"fmt"
	"DemoApp/basicoperations"
)

func main() {
	fmt.Print("Hello Go!")
	testBasicMathOps()
}


func testBasicMathOps(){
	addRes, err := basicoperations.Add(20, 10)
	fmt.Print(" Addtion 20 10 is " , addRes , " Error is " , err)
}