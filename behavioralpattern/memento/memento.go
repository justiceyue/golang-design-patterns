package memento

import "fmt"

/*
在不破坏封装性的前提下，保存一个对象的内部状态
不破坏封装性--->不对外
保存内部状态-->便于恢复到这个状态
用来创建程序某个时刻运行状态的快照
撤销操作的第二种方式:存储恢复式撤销
*/

type Memento interface{}

type ConcreateMemento struct {
	result int64
}

type Command interface {
	Execute()
	Undo(m Memento)
	Redo(m Memento)
	CreateMemento() Memento
}

// 原发器，备忘录保存的是原发器的某个状态
type Opeartion interface {
	GetResult() int64
	Add(num int64)
	Sub(num int64)
	CreateMemento() Memento
	SetMemento(m Memento)
}

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

func (o *Operator) CreateMemento() Memento {
	return &ConcreateMemento{
		result: o.result,
	}
}

func (o *Operator) SetMemento(m Memento) {
	o.result = m.(*ConcreateMemento).result
}

// 利用组合模式完成command统一操作
type commandMeta struct {
	operation Opeartion
}

func (c *commandMeta) CreateMemento() Memento {
	return c.operation.CreateMemento()
}

func (c *commandMeta) Undo(m Memento) {
	c.operation.SetMemento(m)
	fmt.Printf("撤销后的结果为:%d\n", c.operation.GetResult())
}

func (c *commandMeta) Redo(m Memento) {
	c.operation.SetMemento(m)
	fmt.Printf("恢复后的结果为:%d\n", c.operation.GetResult())
}

type AddCommand struct {
	*commandMeta
	num int64
}

//只需各自实现特定操作即可
func (a *AddCommand) Execute() {
	a.operation.Add(a.num)
	fmt.Printf("加后的结果为:%d\n", a.operation.GetResult())
}

type SubCommand struct {
	*commandMeta
	num int64
}

func (s *SubCommand) Execute() {
	s.operation.Sub(s.num)
	fmt.Printf("减后的结果为:%d\n", s.operation.GetResult())
}

//
type Calculator struct {
	addCommand  Command
	subCommand  Command
	undoCommand []Command
	redoCommand []Command
	undoMemento []Memento
	redoMemento []Memento
}

func (c *Calculator) AddPress() {
	//记录之前的状态
	c.undoMemento = append(c.undoMemento, c.addCommand.CreateMemento())
	c.addCommand.Execute()
	c.undoCommand = append(c.undoCommand, c.addCommand)
}

func (c *Calculator) SubPress() {
	//记录之前的状态
	c.undoMemento = append(c.undoMemento, c.subCommand.CreateMemento())
	c.subCommand.Execute()
	c.undoCommand = append(c.undoCommand, c.subCommand)
}

func (c *Calculator) Undo() {
	if len(c.undoCommand) <= 0 {
		fmt.Println("没有可撤销操作")
		return
	}
	currentUndoIndex := len(c.undoCommand) - 1
	c.redoMemento = append(c.redoMemento, c.undoCommand[currentUndoIndex].CreateMemento())
	c.undoCommand[currentUndoIndex].Undo(c.undoMemento[currentUndoIndex])
	c.redoCommand = append(c.redoCommand, c.undoCommand[currentUndoIndex])
	c.undoCommand = c.undoCommand[:currentUndoIndex]
	c.undoMemento = c.undoMemento[:currentUndoIndex]
}

func (c *Calculator) Redo() {
	if len(c.redoCommand) <= 0 {
		fmt.Println("没有可撤销操作")
		return
	}
	currentRedoIndex := len(c.redoCommand) - 1
	c.undoMemento = append(c.undoMemento, c.redoCommand[currentRedoIndex].CreateMemento())
	c.redoCommand[currentRedoIndex].Redo(c.redoMemento[currentRedoIndex])
	c.undoCommand = append(c.undoCommand, c.redoCommand[currentRedoIndex])
	c.redoCommand = c.redoCommand[:currentRedoIndex]
	c.redoMemento = c.redoMemento[:currentRedoIndex]
}
