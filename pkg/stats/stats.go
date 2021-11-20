package stats

import "github.com/Umar2505/bank/v2/pkg/types"

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	var payment types.Payment
	var a int
	var b int
	var c int
	for _, payment = range payments {
		if payment.Category == "auto" {
			categories["auto"] += payment.Amount
			a++
		} else if payment.Category == "food" {
			categories["food"] += payment.Amount
			b++
		} else if payment.Category == "fun" {
			categories["fun"] += payment.Amount
			c++
		}
	}
	categories["auto"] /= types.Money(a)
	categories["food"] /= types.Money(b)
	categories["fun"] /= types.Money(c)
	return categories
}