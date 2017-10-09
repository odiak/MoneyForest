package store

type Account struct {
	ID          string
	OwnerID     string
	Owner       *User
	Name        string
	Description string
	AccountType string
	Balance     int32
	HasBalance  bool
}
