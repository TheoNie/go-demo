package basic

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Log("script_name:" + os.Args[0])

	//循环中拼接字符串，大数据量下效率低下，占用资源,多出来的临时字符串会在适当的时机再垃圾回收
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	t.Log(s)

	//直接通过string库，不产生临时字符串
	t.Log(strings.Join(os.Args[1:], " "))
}

//命令行参数  xxx.go -n xx -s xx xxx xxx xxx
func TestEcho2(t *testing.T) {
	n := flag.Bool("n", false, "it's n")
	sep := flag.String("s", " ", "sep")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}