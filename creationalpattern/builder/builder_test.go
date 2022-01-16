package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	in := &InsuranceContractConcreate{
		InsuranceContract: &InsuranceContract{},
	}
	d := NewInsuranceContractDirector(in)
	err := d.Direct(1)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", in.GetInsuranceContract().PersonName)
	assert.Equal(t, int64(123), in.GetInsuranceContract().StartTime)
}
