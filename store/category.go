package store

type Category struct {
	ID               string `sql:"type:uuid"`
	Name             string
	ParentCategoryID string `sql:"type:uuid"`
	ParentCategory   *Category
}
