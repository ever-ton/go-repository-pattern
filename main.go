package main

import (
	"database/sql"
	"fmt"

	"github.com/ever-ton/go-repository-pattern/repository"
	"github.com/ever-ton/go-repository-pattern/service"
)

func main() {

	// db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	db, _ := sql.Open("sqLite3", "sqlite.repository")
	repositorySqLite := repository.NewCategorySqlite(db)

	service := service.NewCategoryService(repositorySqLite)
	cat, _ := service.Create("Minha Categoria")

	fmt.Println(cat)

}
