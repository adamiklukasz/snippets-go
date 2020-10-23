package databases

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	fmt.Printf("db=%#v\n", db)
	fmt.Printf("err=%#v\n", err)

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	rp := Product{}
	db.Update("price", 300).Where("code = ?", "D42")
	db.First(&rp, "code = ?", "D42")

	fmt.Printf("rp=%v\n", rp.Price)
}
