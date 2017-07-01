package basicoperations

func Add(a int, b int) (result int, err error){
	return a + b, nil;
}

func Subtract(a int, b int) (result int, err error){
	return a - b, nil
}

func Multiply(a int, b int) (result int, err error){
	return a * b, nil
}

func Divide(a int, b int) (result int, err error){
	return a / b, nil;
}
