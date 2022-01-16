package observer

import "fmt"

/*
处理一对多关系
可以将中介者看成是负责的观察者模式
*/

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	notifyAll()
}

//一个商品维护一组观察者----建立关系
type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o Observer) {
	var tempObverserList []Observer
	for _, v := range i.observerList {
		if v.GetID() != o.GetID() {
			tempObverserList = append(tempObverserList, v)
		}
	}
	i.observerList = tempObverserList
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

type Observer interface {
	Update(string)
	GetID() string
}

type customer struct {
	id string
}

func (c *customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) GetID() string {
	return c.id
}
