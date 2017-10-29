package store

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Category struct {
	ID               string `sql:"type:uuid"`
	OwnerID          string `sql:"type:uuid,notnull"`
	Owner            *User
	Name             string  `sql:",notnull"`
	ParentCategoryID *string `sql:"type:uuid"`
	ParentCategory   *Category
	ChildCategories  []Category `pg:",fk:ParentCategory"`
}

func (c *Category) validate(db orm.DB) error {
	if c.Name == "" {
		return ValidationError("name is required")
	}
	if c.ParentCategoryID != nil {
		parent := Category{}
		err := db.Model(&parent).Where("owner_id = ? AND id = ?", c.OwnerID, *c.ParentCategoryID).Limit(1).Select()
		if err != nil {
			if err == pg.ErrNoRows {
				return ValidationError("parent-category-not-found")
			}
			return err
		}
		if parent.ParentCategoryID != nil {
			return ValidationError("parent-has-parent")
		}
	}

	return nil
}

func (c *Category) BeforeInsert(db orm.DB) error {
	return c.validate(db)
}

func (c *Category) BeforeUpdate(db orm.DB) error {
	return c.validate(db)
}
