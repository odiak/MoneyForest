package store

type Category struct {
	ID               string `sql:"type:uuid"`
	OwnerID          string `sql:"type:uuid"`
	Owner            *User
	Name             string
	ParentCategoryID string `sql:"type:uuid"`
	ParentCategory   *Category
}
