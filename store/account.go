package store

type Account struct {
	ID          string `sql:"type:uuid"`
	OwnerID     string `sql:"type:uuid"`
	Owner       *User
	Name        string
	Description string
	AccountType string
	Balance     int32
	HasBalance  bool
}
