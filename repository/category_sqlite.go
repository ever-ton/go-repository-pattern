package repository

import (
	"database/sql"

	"github.com/ever-ton/go-repository-pattern/entity"
	_ "github.com/mattn/go-sqlite3"
)

type CategoryRepositorySqlite struct {
	db *sql.DB
}

func NewCategorySqlite(db *sql.DB) *CategoryRepositorySqlite {
	return &CategoryRepositorySqlite{db: db}
}

func (c *CategoryRepositorySqlite) Get(id string) (entity.Category, error) {
	var category entity.Category
	stmt, err := c.db.Prepare("select id, name from categories where id=?")
	if err != nil {
		return entity.Category{}, err
	}
	err = stmt.QueryRow(id).Scan(&category.ID, &category.Name)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (c *CategoryRepositorySqlite) Create(category entity.Category) (entity.Category, error) {
	stmt, err := c.db.Prepare(`insert into categories (id, name) values(?, ?)`)
	if err != nil {
		return entity.Category{}, err
	}

	_, err = stmt.Exec(category.ID, category.Name)
	if err != nil {
		return entity.Category{}, err
	}

	err = stmt.Close()
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}
