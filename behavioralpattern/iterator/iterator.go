package iterator

/*
command和mediator中有用到go两种常用的iterator模式

command:
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

**该种为经典实现，有涉及index逻辑的时候可以选用这种

mediator:
type dataIteraor map[string][]string

func (d dataIteraor) deleteKey(fn func(name string) bool) {
	for key := range d {
		if fn(key) {
			delete(d, key)
		}
	}
}

**go中比较经典的涉及，很多sdk中都有体现，如标准库中container/ring

*/
