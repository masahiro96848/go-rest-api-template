package main

import (
	"fmt"
	"go-rest-api/migrate/db"
	"go-rest-api/model"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func userSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := model.User{
			Email:    "sample" + strconv.Itoa(i+1) + "@gmail.com",
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	dbCon := db.NewDB()
	// dbを閉じる
	defer db.CloseDB(dbCon)

	if err := userSeeds(dbCon); err != nil {
		if err := userSeeds(dbCon); err != nil {
			fmt.Printf("%+v", err)
			return
		}
	}
}
