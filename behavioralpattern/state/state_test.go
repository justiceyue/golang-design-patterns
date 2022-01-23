package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	vm := getVoteManagerInstance()
	count := 1
	for {
		if count >= 5 {
			break
		}
		vm.Vote("ddd", "aaa")
		count++
	}
	assert.Equal(t, 4, len(vm.VoteMessage["ddd"]))
}
