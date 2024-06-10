package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) (rate float32) {
	if balance < 0 {
		rate = 3.213
	} else if balance >= 0 && balance < 1000 {
		rate = 0.500000
	} else if balance >= 1000 && balance < 5000 {
		rate = 1.621
	} else {
		rate = 2.475
	}
	return
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) (interest_balance float64) {
	rate := float64(InterestRate(balance) / 100)
	interest_balance = balance * (rate)
	return
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) (balanced_updated float64) {
	balanced_updated = Interest(balance) + balance
	return
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) (years int) {
	expected_balance := balance
	years = 0
	for expected_balance < targetBalance {
		expected_balance += Interest(balance)
		balance = AnnualBalanceUpdate(balance)
		years += 1
	}
	return
}
