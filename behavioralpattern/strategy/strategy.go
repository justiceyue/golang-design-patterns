package strategy

import "fmt"

/*
	具体的算法实现从业务中分离出来，这些算法可以相互替换，但是每一时刻只能有一个算法被使用
	如何组织和调用这些算法是关键
*/

type LogStrategy interface {
	Log(msg string)
}

type LogFunc func(msg string)

func (l LogFunc) Log(msg string) {
	l(msg)
}

func DBLog(msg string) {
	if len(msg) > 5 {
		panic(msg)
	}
	fmt.Printf("将日志%s记录到db中", msg)
}

func FileLog(msg string) {
	fmt.Printf("将日志%s记录到文件中", msg)
}

type LogContext struct {
	logStrategy LogStrategy
}
