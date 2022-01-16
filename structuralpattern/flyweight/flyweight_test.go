package flyweight

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	CT := &player{
		identity: "CT",
	}
	T := &player{
		identity: "T",
	}
	getter := arsenalGetter()
	CT.arsenal = getter(CT.identity)
	T.arsenal = getter(T.identity)

	assert.Equal(t, true, reflect.DeepEqual(CT.arsenal.gun, ctGun))
	assert.Equal(t, true, reflect.DeepEqual(T.arsenal.gun, tGun))
}
