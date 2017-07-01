package addition

import "fmt"
import "runtime"

func Add(a int, b int) (result int, err error)  {
	fmt.Println("Called add for ", a, b)
	return (a + b),nil;
}

func PrintNoOfCores () {
	cores := runtime.NumCPU()
	fmt.Println("This machine has ", cores, " cores")
}