package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

//通道类型的值本身就是并发安全的，这也是 Go 语言自带的、唯一一个可以满足并发安全性的类型。
//一个通道相当于一个先进先出（FIFO）的队列。
//通道中的各个元素值都是严格地按照发送的顺序排列的，先被发送通道的元素值一定会先被接收。
func TestChannel(t *testing.T) {
	ch := make(chan string, 3) //make是go用来做初始化的函数。chan是通道类型的关键词，chan string代表一个传递string类型值的通道，3代表通道的容量，最多有3个，如果是0的话，这个属于无缓冲通道。
	ch <- "hello"
	ch <- "world"
	ch <- "haha"

	msg := <-ch
	t.Log(msg)
}

//定义了一个返回值是接收通道的函数，只能从该通道读取值
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

//除非有特殊的保障措施，否则千万不要让接收方去关闭通道，而应该让发送方去做这件事。
func TestSingleChannel(t *testing.T) {
	intChan2 := getIntChan()
	for elem := range intChan2 { //因为上面调用的函数中的通道已经关闭，所以这里不会产生阻塞，只会利用range把通道中的所有值读取出来
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
}

//在使用select语句的时候，如果设置有默认分支，则即使其他case有阻塞，select也不会阻塞；如果没有默认分支，那么所有case都不满足条件的时候就会出现阻塞，直到其中任何一个case条件被满足。
func TestSelectChannel(t *testing.T) {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}
