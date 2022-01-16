package abstractfactory

// cpu
type CPU interface {
	Stitch() int64
	Brand() string
}

type intel struct {
	// 针脚
	stitch int64
}

func (i intel) Stitch() int64 {
	return i.stitch
}

func (i intel) Brand() string {
	return "intel"
}

type amd struct {
	// 针脚
	stitch int64
}

func (a amd) Stitch() int64 {
	return a.stitch
}

func (a amd) Brand() string {
	return "amd"
}

// mainboard
type Mainboard interface {
	// 安装cpu true-成功 false-失败
	InstallCPU(cpu CPU) bool
	GetInstalledCPU() string
}

type msi struct {
	// 插槽
	slot int64
	// cpu
	cpu string
}

func (m *msi) InstallCPU(cpu CPU) bool {
	if m.slot == cpu.Stitch() {
		m.cpu = cpu.Brand()
		return true
	}
	return false
}

func (m *msi) GetInstalledCPU() string {
	return m.cpu
}

type gam struct {
	//插槽
	slot int64
	cpu  string
}

func (g *gam) InstallCPU(cpu CPU) bool {
	if g.slot == cpu.Stitch() {
		g.cpu = cpu.Brand()
		return true
	}
	return false
}

func (g *gam) GetInstalledCPU() string {
	return g.cpu
}

// 如果单纯的使用工厂模式分别创建cpu和mainboard,则很容易导致cpu针脚和主板槽位不匹配
/*
type CPUFactory interface {
	CreateCPU(brand string, stitch int64) CPU
}

type MainBoardFactory interface {
	CreateMainBoard(brand string, slot int64) Mainboard
}

type cpuFactory struct{}

// 参数化工厂
func (cpuFactory) CreateCPU(brand string, stitch int64) CPU {
	if brand == "intel" {
		return intel{
			stitch: stitch,
		}
	}
	return amd{
		stitch: stitch,
	}
}

type mainBoardFactory struct{}

func (mainBoardFactory) CreateMainBoard(brand string, slot int64) Mainboard {
	if brand == "msi" {
		return msi{
			slot: slot,
		}
	}
	return gam{slot: slot}
}
*/

type PCComponentsWorkShop interface {
	CreateCPU(brand string) CPU
	CreateMainBoard(brand string) Mainboard
	AssembleMainboardWithCPU(cpu CPU, mainboard Mainboard) bool
}

type AssemblyLineA struct {
	Stitch int64
	Slot   int64
}

func (a AssemblyLineA) CreateCPU(brand string) CPU {
	if brand == "intel" {
		return intel{stitch: a.Stitch}
	}
	return amd{stitch: a.Stitch}
}

func (a AssemblyLineA) CreateMainBoard(brand string) Mainboard {
	if brand == "msi" {
		return &msi{slot: a.Slot}
	}
	return &gam{slot: a.Slot}
}

func (AssemblyLineA) AssembleMainboardWithCPU(cpu CPU, mainboard Mainboard) bool {
	return mainboard.InstallCPU(cpu)
}
