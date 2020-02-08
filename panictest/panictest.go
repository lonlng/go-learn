package panictest

import "fmt"

func Test_01() {

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		println(err.(string)) // 将 interface{} 转型为具体类型。
	// 	}
	// }()
	defer func() {
		fmt.Println("defer !")

	}()
	panic("panic")
}
