package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactory(t *testing.T) {
	assemblyLineA := AssemblyLineA{
		Stitch: 164,
		Slot:   164,
	}
	cpuBrand := "intel"
	mainboardBrand := "msi"
	mainboard, ok := AssembleMainboard(assemblyLineA, cpuBrand, mainboardBrand)
	assert.Equal(t, true, ok)
	assert.Equal(t, cpuBrand, mainboard.GetInstalledCPU())
}

func AssembleMainboard(pcw PCComponentsWorkShop, cpuBrand, mainBoardBrand string) (Mainboard, bool) {
	cpu := pcw.CreateCPU(cpuBrand)
	mainboard := pcw.CreateMainBoard(mainBoardBrand)
	status := pcw.AssembleMainboardWithCPU(cpu, mainboard)
	return mainboard, status
}
