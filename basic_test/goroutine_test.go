package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	//异步执行，所以下面的语句可以直接打印而不会等待
	go func() {
		time.Sleep(10 * time.Second)
	}()
	t.Log("hahahahaha")
}

func TestGoroutine2(t *testing.T) {
	go spinner(100 * time.Millisecond, t)
	const n = 45
	fibN := fib(n)  //slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration, t *testing.T) {
	for  {
		for _, r := range `.。`{
			t.Logf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
