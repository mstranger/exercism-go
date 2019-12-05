package account

import "sync"

// Account represents a bank account.
type Account struct {
	amount int64
	active bool
	mux    sync.Mutex
}

// Open creates a new Account with innitial deposit.
func Open(initial int64) *Account {
	if initial < 0 {
		return nil
	}

	return &Account{
		amount: initial,
		active: true,
	}
}

// Close closes the account and returns the balance.
func (acc *Account) Close() (int64, bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	// already closed
	if !acc.active {
		return acc.amount, false
	}

	payout := acc.amount
	acc.active = false
	acc.amount = 0

	return payout, true
}

// Balance returns the current amount of money.
func (acc *Account) Balance() (int64, bool) {
	if !acc.active {
		return 0, false
	}

	return acc.amount, true
}

// Deposit adds/withdrawal amount to/from current balance.
func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	// account closed
	if !acc.active {
		return 0, false
	}

	newAmount := acc.amount + amount
	if newAmount < 0 {
		return newAmount, false
	}

	acc.amount = newAmount
	return newAmount, true
}
