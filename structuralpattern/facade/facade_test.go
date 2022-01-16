package facade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	assert.NoError(t, err)

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	assert.NoError(t, err)
}
