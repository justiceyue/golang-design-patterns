package simplefactory

import "fmt"

type Car interface {
	Dirve()
}

type bmw struct {
}

func (bmw) Dirve() {
	fmt.Println("bmw")
}

type tesla struct {
}

func ProduceCar(brand string) Car {
	if brand == "tesla" {
		return tesla{}
	}
	return bmw{}
}

func (tesla) Dirve() {
	fmt.Println("tesla")
}
