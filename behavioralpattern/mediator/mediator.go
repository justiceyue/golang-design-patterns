package mediator

import "sync"

/*
使用中介者模式让多对多关系简化为两个一对多关系
部门和工作人员之间是多对多关系
当撤销，合并部门和人员离职，调岗的时候，多对多关系中删除操作会变的复杂
*/

// test data
// 模拟部门表和职工表
var (
	DepData map[string][]string = map[string][]string{
		"personnel":  {"jack", "rose", "tonny"},
		"technology": {"jack"},
		"accounting": {"tonny"},
	}

	StaffData map[string][]string = map[string][]string{
		"jack":  {"personnel", "technology"},
		"rose":  {"personnel"},
		"tonny": {"accounting", "personnel"},
	}

	once     sync.Once
	mediator *Mediator
)

//部门
type Dep struct {
	ID   int64
	Name string
}

func (d *Dep) Cancel() {
	// 让中介者去真正撤销部门
	getMediatorInstance().DeleteDep(d.Name)
}

type Staff struct {
	ID   int64
	Name string
}

func (s *Staff) Dismiss() {
	// 让中介者去真正执行人员离职操作
	getMediatorInstance().DeleteStaff(s.Name)
}

//需维护关系
//可将dep作为一个接口，每个部门单独实现，部门中有职员的list即可,职员同部门实现。
type Mediator struct {
	Dep   dataIteraor
	Staff dataIteraor
}

// 使用单例模式获取mediator
func getMediatorInstance() *Mediator {
	once.Do(func() {
		mediator = &Mediator{
			Dep:   DepData,
			Staff: StaffData,
		}
	})
	return mediator
}

func (m *Mediator) DeleteDep(name string) {
	// 先删除staff中的dep
	m.Staff.deleteValue(func(_name string) bool {
		return _name == name
	})
	//删除dep
	m.Dep.deleteKey(func(_name string) bool {
		return _name == name
	})
}

func (m *Mediator) DeleteStaff(name string) {
	m.Dep.deleteValue(func(_name string) bool {
		return _name == name
	})
	m.Staff.deleteKey(func(_name string) bool {
		return _name == name
	})
}

type dataIteraor map[string][]string

func (d dataIteraor) deleteKey(fn func(name string) bool) {
	for key := range d {
		if fn(key) {
			delete(d, key)
		}
	}
}

func (d dataIteraor) deleteValue(fn func(name string) bool) {
	for key, values := range d {
		var tempValue []string
		for _, value := range values {
			if !fn(value) {
				tempValue = append(tempValue, value)
			}
		}
		d[key] = tempValue
	}
}
