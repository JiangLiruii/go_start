package main
import (
	"flag"
	"fmt"
)
// 方法1
var name = "everyone";
func init()  {
	// 存储name 的地址 命令参数的名称 制定未追加改命令参数时的默认值 简短说明
	flag.StringVar(&name, "name1", "everyone", "The greeting Object")
	flag.StringVar(&name, "name2", "Not anyone", "The greeting Object")
}

func main()  {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

// go run demo1.go -name2=Jose --> Hello, Jose!
// go run demo1.go -name1=Jose --> Hello, Jose! (strange)
// go run demo1.go -name1=Jose -name2=Mike --> Hello, Mike!
// go run demo1.go -name2=Jose -name1=Mike --> Hello, Mike! (strange)

