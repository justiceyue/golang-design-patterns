package chain

import "fmt"

/*
一个请求需要被多个对象进行处理，可以将处理对象连成一条链，挨个处理请求。
传统的职责链一般只有一个对象可以对请求进行处理，实际开发过程中使用比较多的是功能链。
模仿Gin，写一个全局中间件
*/

type HandlerFunc func()

type HandlerFuncChain []HandlerFunc

type RouterGroup struct {
	Handlers HandlerFuncChain
	index    int
}

func (r *RouterGroup) Use(handlers ...HandlerFunc) {
	r.Handlers = append(r.Handlers, handlers...)
}

func (r *RouterGroup) Next() {
	for r.index < len(r.Handlers) {
		r.Handlers[r.index]()
		r.index++
	}
}

func middleware1() {
	fmt.Println("中间件1执行完毕")
}

func middleware2() {
	fmt.Println("中间件2执行完毕")
}
