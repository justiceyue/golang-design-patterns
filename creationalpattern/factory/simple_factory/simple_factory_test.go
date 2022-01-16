package simplefactory

import (
	"testing"
)

func TestCar(t *testing.T) {
	tesla := ProduceCar("tesla")
	tesla.Dirve()
}
