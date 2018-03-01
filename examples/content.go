package main

import (
	"context"
	"fmt"
	"time"
)

//http://www.flysnow.org/2017/05/12/go-in-action-go-context.html  比较好解释关于context机制的介绍。
//Context 使用原则
//不要把Context放在结构体中，要以参数的方式传递
//以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
//给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
//Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
//Context是县城安全的，可以放心的在多个goroutine中传递

func main() {
	//context1()
	//context2()

	//context3()
	ExampleWithTimeout()

}

func ExampleWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Output:
	// context deadline exceeded
}

func context3() {
	ctx, cancel := context.WithCancel(context.Background())
	//同一个context可以控制多个goroutine的生命周期。
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func context2() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx) //当作参数传递给goroutine,并且必须以参数的形式传去进去。不能作为一个对象

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel() //调用的生成对应的配对的取消函数。
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func context1() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
