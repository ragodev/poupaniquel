package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFromRawToTransactionAndAddChildrenRecursively(t *testing.T) {
	raw := []Raw{
		Raw{Id: 1, Description: "Mastercard", ParentId: 0, RecursiveNumber: 1},
		Raw{Id: 3, Description: "Clothes", ParentId: 0, RecursiveNumber: 1},
		Raw{Id: 4, Description: "Internet", ParentId: 0, RecursiveNumber: 1},
		Raw{Id: 2, Description: "Superstore", ParentId: 1, RecursiveNumber: 2},
		Raw{Id: 8, Description: "Travel", ParentId: 1, RecursiveNumber: 2},
		Raw{Id: 6, Description: "Bread", ParentId: 2, RecursiveNumber: 3},
		Raw{Id: 7, Description: "Drink", ParentId: 2, RecursiveNumber: 3},
		Raw{Id: 9, Description: "Flight to Brazil", ParentId: 8, RecursiveNumber: 3},
		Raw{Id: 5, Description: "Soda", ParentId: 7, RecursiveNumber: 4},
	}

	trans := toTransaction(raw)
	assert.Equal(t, 3, len(trans))

	mastercard := trans[0]
	assert.Equal(t, int64(1), mastercard.Id)
	assert.Equal(t, "Mastercard", mastercard.Description)
	assert.Equal(t, 2, len(mastercard.Children))

	superstore := mastercard.Children[1]
	assert.Equal(t, int64(2), superstore.Id)
	assert.Equal(t, "Superstore", superstore.Description)
	assert.Equal(t, 2, len(superstore.Children))

	drink := superstore.Children[0]
	assert.Equal(t, int64(7), drink.Id)
	assert.Equal(t, "Drink", drink.Description)
	assert.Equal(t, 1, len(drink.Children))

	soda := drink.Children[0]
	assert.Equal(t, int64(5), soda.Id)
	assert.Equal(t, "Soda", soda.Description)
	assert.Equal(t, 0, len(soda.Children))
}
