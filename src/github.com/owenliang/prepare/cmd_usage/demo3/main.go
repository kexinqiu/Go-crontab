package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

//结构体
type result struct {
	err    error
	output []byte
}

func main() {
	//  执行1个cmd, 让它在一个协程里去执行, 让它执行2秒: sleep 2; echo hello;
	// 1秒的时候, 我们杀死cmd
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// 创建了一个结果队列
	resultChan = make(chan *result, 1000)

	// context:   chan byte
	// cancelFunc:  close(chan byte)
	//context ctx 里面有一个chanel， chanel有一个取消函数，会关掉chanel， ctx会感知到chanel是否被关闭

	//杀死程序   context继承了context.TODO
	//创建了一个上下文和一个取消函数
	ctx, cancelFunc = context.WithCancel(context.TODO())

	//协程
	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext( /* context 取消命令执行*/ ctx, "/bin/bash", "-c", "sleep 2;echo hello;")

		// 执行任务, 捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果, 传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	//主程序不会等待协程，执行1秒就被终止
	// 继续往下走
	time.Sleep(1 * time.Second)

	// 取消上下文
	cancelFunc()

	// 在main协程里, 等待子协程的退出，并打印任务执行结果
	res = <-resultChan

	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
