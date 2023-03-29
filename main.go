package main

import "fmt"

func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}
func main() {
	sum, diff := sumAndDiff(2, 2)
	fmt.Printf("%v %v", sum, diff)
}
