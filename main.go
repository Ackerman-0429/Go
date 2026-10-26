// main.go —— Go 开发环境自检示例
//
// 运行方式：
//   go run .
//   go run main.go
// 编译成可执行文件：
//   go build -o hello.exe .
package main

import (
	"fmt"
	"runtime"
)

// Greeter 一个简单的结构体，演示 Go 的"方法"写法
type Greeter struct {
	Name string
}

// Hello 是挂在 Greeter 上的方法
func (g Greeter) Hello() string {
	return fmt.Sprintf("你好，%s！", g.Name)
}

func main() {
	fmt.Println("==============================")
	fmt.Println("  Go 开发环境配置成功")
	fmt.Println("==============================")
	fmt.Println(Greeter{Name: "Go 开发者"}.Hello())
	fmt.Printf("Go 版本  : %s\n", runtime.Version())
	fmt.Printf("操作系统 : %s / %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("CPU 核数 : %d\n", runtime.NumCPU())
	fmt.Println("------------------------------")

	// 简单的循环，验证基本语法
	sum := 0
	for i := 1; i <= 3; i++ {
		fmt.Printf("循环输出 : 第 %d 次\n", i)
		sum += i
	}
	fmt.Printf("1 + 2 + 3 = %d\n", sum)
	fmt.Println("==============================")
}
