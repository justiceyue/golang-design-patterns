package memento

import "testing"

func TestMemento(t *testing.T) {
	op := &Operator{
		result: 1,
	}
	addCommand := &AddCommand{
		commandMeta: &commandMeta{
			operation: op,
		},
		num: 1,
	}
	subCommand := &SubCommand{
		commandMeta: &commandMeta{
			operation: op,
		},
		num: 1,
	}
	cal := &Calculator{
		addCommand: addCommand,
		subCommand: subCommand,
	}
	cal.AddPress() //2
	cal.AddPress() //3
	cal.Undo()     //2
	cal.Redo()     //3
	cal.Undo()     //2
	cal.SubPress() //1
	cal.Undo()     //2
	cal.Undo()     //1
	cal.Undo()     //无可撤销操作
}
