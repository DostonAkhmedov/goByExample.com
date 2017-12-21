package main

import "fmt"

func intSeq() func() int {
	var i int = 0
	return func() int {
		i ++;
		return i
	}
}

func main() {

	nextInt := intSeq()


	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	var newInts = intSeq()
	fmt.Println(newInts())

}