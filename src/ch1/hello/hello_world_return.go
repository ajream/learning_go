package main //包，表明代码所在的模块（包）

//引入代码依赖
import (
	"fmt"
	"os"
)

// 功能实现，main 函数不支持传入参数
// 在程序中直接通过 os.Args 获取命令行参数
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello ", os.Args[1])
	}
}
