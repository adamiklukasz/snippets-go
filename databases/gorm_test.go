package databases

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func TestGORM(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()

	db.AutoMigrate(&Product{})

	p1 := &Product{
		Code:  "L121",
		Price: 1000,
	}
	db.Create(p1)

	var pres Product
	db.Find(&pres, "code = ?", "L121")

	db.Model(&pres).Update("Price", 2000)
}
