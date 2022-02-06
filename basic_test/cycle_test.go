package basic

import (
	"fmt"
	"testing"
)

//循环
func TestCycle(t *testing.T) {
	n := 0
	for n <= 3 {
		n++
		fmt.Println(n)
	}

	for i := 1; i <= 3; i++ {
		//...
	}
}

//死循环
func TestEndlessLoop(t *testing.T) {
	for {
		fmt.Println("not end")
	}
}
