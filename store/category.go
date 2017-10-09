package store

type Category struct {
	ID               string
	OwnerID          string
	Owner            *User
	Name             string
	ParentCategoryID string
	ParentCategory   *Category
}
