package adaptor

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type mac struct {
}

func (mac) InsertIntoLightningPort() {
	fmt.Println("InsertIntoLightningPort success!")
}

type windows struct {
}

func (windows) InsertIntoUSBPort() {
	fmt.Println("InsertIntoUSBPort success!")
}

type windowsLightningAdaptor struct {
	w windows
}

func (wla windowsLightningAdaptor) InsertIntoLightningPort() {
	wla.w.InsertIntoUSBPort()
	fmt.Println("adaptor Is to take effect")
	fmt.Println("InsertIntoLightningPort success!")
}
