package command

import "testing"

func TestCommand(t *testing.T) {
	op := &Operator{}
	addCommand := &AddCommand{
		Operation: op,
		Number:    1,
	}
	subCommand := &SubCommand{
		Operation: op,
		Number:    1,
	}
	cal := &Calculator{
		AddCommand: addCommand,
		SubCommand: subCommand,
	}
	cal.AddPress()   //1
	cal.AddPress()   //2
	cal.UndoPress(1) //1
	cal.AddPress()   //2
	cal.RedoPress(1) //3
}
