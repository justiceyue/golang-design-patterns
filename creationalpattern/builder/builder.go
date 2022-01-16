package builder

import "errors"

/*
 创建一份保险合同
 具体约束如下:
 1.保险合同要么和个人，要么和公司签订
 2.开始时间小于结束时间
*/

// 产品---保险合同
type InsuranceContract struct {
	ID          string
	PersonName  string
	CompanyName string
	StartTime   int64
	EndTime     int64
}

// 保险合同建造者抽象
type InsuranceContractBuilder interface {
	BuildPerson(name string)
	BuildCompanyName(name string)
	AddCommencementDate(startTime, endTime int64)
	Validate() error
	GetInsuranceContract() *InsuranceContract
}

type InsuranceContractConcreate struct {
	*InsuranceContract
}

func (i *InsuranceContractConcreate) BuildPerson(name string) {
	i.PersonName = name
}

func (i *InsuranceContractConcreate) BuildCompanyName(name string) {
	i.CompanyName = name
}

func (i *InsuranceContractConcreate) AddCommencementDate(startTime, endTime int64) {
	i.StartTime = startTime
	i.EndTime = endTime
}

func (i *InsuranceContractConcreate) Validate() error {
	if i.CompanyName != "" && i.PersonName != "" {
		return errors.New("too many names")
	}
	if i.StartTime > i.EndTime {
		return errors.New("time error")
	}
	return nil
}

func (i *InsuranceContractConcreate) GetInsuranceContract() *InsuranceContract {
	return i.InsuranceContract
}

// 合同签写指导
type InsuranceContractDirector struct {
	ICB InsuranceContractBuilder
}

func NewInsuranceContractDirector(builder InsuranceContractBuilder) *InsuranceContractDirector {
	return &InsuranceContractDirector{
		ICB: builder,
	}
}

func (ICD *InsuranceContractDirector) Direct(signObj int64) error {
	if signObj == 1 {
		ICD.ICB.BuildPerson("TEST")
	} else {
		ICD.ICB.BuildCompanyName("TEST_CO")
	}
	ICD.ICB.AddCommencementDate(123, 345)
	if err := ICD.ICB.Validate(); err != nil {
		return err
	}
	return nil
}
