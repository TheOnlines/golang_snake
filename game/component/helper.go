package component

import (
	"fmt"
	"time"
)

// Catch 全局捕获
func Catch() {
	if err := recover(); err != nil {

		//打印异常
		fmt.Println(fmt.Sprintf("Catch faild:%s", err))

		//获取堆栈信息(暂时先不获取吧.)
		/*		buf := new(bytes.Buffer)
				fmt.Fprintf(buf, "%v\n", err)
				for i := 1; ; i++ {
					pc, file, line, ok := getControl.Caller(i)
					if !ok {
						break
					}
					fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
				}
				fmt.Println(buf.String())*/
	}
}

// Run 执行方法
func Run(f func()) {
	defer Catch()
	f()
}

//flush 刷新
func flush(score int) {
	time.Sleep(time.Duration(150-(score/10)) * time.Millisecond)
}
