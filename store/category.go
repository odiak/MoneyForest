package store

type Category struct {
	ID               string `sql:"type:uuid,notnull"`
	OwnerID          string `sql:"type:uuid,notnull"`
	Owner            *User
	Name             string `sql:",notnull"`
	ParentCategoryID string `sql:"type:uuid,notnull"`
	ParentCategory   *Category
}
