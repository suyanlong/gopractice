package goroutine

import (
	"fmt"
)

func goroutine() {
	//quit := make(chan int)
	//quit <- 2   无缓存的通道是一直阻塞的。
	fmt.Println("end")
	//<-quit

}
