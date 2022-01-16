package observer

import "testing"

func TestObverser(t *testing.T) {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()

	shirtItem.deregister(observerFirst)
	shirtItem.updateAvailability()
}
