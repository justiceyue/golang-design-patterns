package mediator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediator(t *testing.T) {
	//人事部实例
	//这里可用简单工厂创建实例
	p := &Dep{
		Name: "personnel",
	}
	p.Cancel()
	_, ok := mediator.Dep[p.Name]
	assert.Equal(t, false, ok)
	for _, deps := range mediator.Staff {
		for _, dep := range deps {
			//没有人事部
			t.Log(dep)
		}
	}
	j := &Staff{
		Name: "jack",
	}
	j.Dismiss()
	_, ok = mediator.Staff[j.Name]
	assert.Equal(t, false, ok)
	for _, staffs := range mediator.Dep {
		for _, staff := range staffs {
			//没有jack,rose(人事部删除了)
			t.Log(staff)
		}
	}
}
