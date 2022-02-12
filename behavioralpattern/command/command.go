package command

import (
	"fmt"
)

/*
请求封装成对象，统一执行操作的接口，可围绕命令对象完成存储，撤销，队列等操作
以可撤销操作为例（反操作式撤销）
*/

// 具体执行操作的接口
type Operation interface {
	GetResult() int64
	Add(num int64)
	Sub(num int64)
}

// 具体实现
type Operator struct {
	result int64
}

func (o *Operator) GetResult() int64 {
	return o.result
}

func (o *Operator) Add(num int64) {
	o.result += num
}

func (o *Operator) Sub(num int64) {
	o.result -= num
}

// 命令接口
type Command interface {
	Execute()
	Undo()
}

// 具体的命令
type AddCommand struct {
	Operation Operation
	Number    int64
}

func (a *AddCommand) Execute() {
	a.Operation.Add(a.Number)
	fmt.Printf("加之后的结果是:%d\n", a.Operation.GetResult())
}

func (a *AddCommand) Undo() {
	a.Operation.Sub(a.Number)
	fmt.Printf("减之后的结果是:%d\n", a.Operation.GetResult())
}

// 具体的命令
type SubCommand struct {
	Operation Operation
	Number    int64
}

func (s *SubCommand) Undo() {
	s.Operation.Add(s.Number)
}

func (s *SubCommand) Execute() {
	s.Operation.Sub(s.Number)
}

//invoker
type Calculator struct {
	AddCommand Command
	SubCommand Command
	//撤销时使用
	UndoList Commands
	//恢复时使用
	RedoList Commands
}

type Commands []Command

func (cs Commands) Iterator() func() (Command, bool) {
	return func() (Command, bool) {
		firstCommandIdx := len(cs) - 1
		if firstCommandIdx < 0 {
			fmt.Println("step过大，没有可撤销的操作了")
			return nil, false
		}
		command := cs[firstCommandIdx]
		cs = cs[:firstCommandIdx] //当[:0]时就是清空切片
		return command, true
	}
}

//此处可体现请求封装成对象
func (c *Calculator) AddPress() {
	c.AddCommand.Execute()
	c.UndoList = append(c.UndoList, c.AddCommand)
}

func (c *Calculator) SubPress() {
	c.SubCommand.Execute()
	c.UndoList = append(c.UndoList, c.SubCommand)
}

func (c *Calculator) UndoPress(step int64) {
	if len(c.UndoList) == 0 {
		fmt.Println("没有可撤销操作")
		return
	}
	iter := c.UndoList.Iterator()
	iterCount := 1
	for {
		if iterCount > int(step) {
			break
		}
		command, hasNext := iter()
		if !hasNext {
			break
		}
		command.Undo()
		c.RedoList = append(c.RedoList, command)
		iterCount++
	}
}

func (c *Calculator) RedoPress(step int64) {
	if len(c.RedoList) == 0 {
		fmt.Println("没有可恢复操作")
		return
	}
	iter := c.RedoList.Iterator()
	iterCount := 1
	for {
		if iterCount > int(step) {
			break
		}
		command, hasNext := iter()
		if !hasNext {
			break
		}
		command.Execute()
		c.UndoList = append(c.UndoList, command)
		iterCount++
	}
}
