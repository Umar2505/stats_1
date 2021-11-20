package stats

import (
	"reflect"
	"testing"
	"github.com/Umar2505/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       123,
			Amount:   50_000_00,
			Category: "auto",
			Status:   "OK",
		},
		{
			ID:       321,
			Amount:   50_000_00,
			Category: "food",
			Status:   "OK",
		},
		{
			ID:       132,
			Amount:   50_000_00,
			Category: "fun",
			Status:   "OK",
		},
	}
	result := CategoriesAvg(payments)
	should := map[types.Category]types.Money{
		"auto" : 5000000,
		"food" : 5000000,
		"fun" : 5000000,
	}

	if !reflect.DeepEqual(result,should) {
		t.Errorf("Didn't true!")
	}
}