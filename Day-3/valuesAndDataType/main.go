package main

import "fmt"

func main() {

	//Non Pointer
	var num int
	num = 32
	num = 25622132

	//pointer
	var addressOfNum *int
	addressOfNum = &num

	//Non Pointer
	var decNum float32
	decNum = 13211.366

	//pointer
	var addressOfdec *float32
	addressOfdec = &decNum

	//Non Pointer
	var flags bool
	flags = true
	var addresOfflags *bool
	addresOfflags = &flags

	//Non Pointer
	var name string
	name = "GoLang By SK"
	var addresOfname *string
	addresOfname = &name

	fmt.Printf("Number value = %d , decimal value = %f \n flags value = %v \n name value = %s \n", num, decNum, flags, name)
	fmt.Printf("addressOfNum value = %v , addressOfdec value = %v \n addresOfflags value = %v \n addresOfname value = %v \n", addresOfname, addressOfdec, addresOfflags, addressOfNum)

}
