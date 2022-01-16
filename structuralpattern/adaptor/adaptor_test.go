package adaptor

import "testing"

func TestAdaptor(t *testing.T)  {
	m := mac{}
	m.InsertIntoLightningPort()

	w := windows{}
	w.InsertIntoUSBPort()

	wla := windowsLightningAdaptor{w: w}
	wla.InsertIntoLightningPort()
}
