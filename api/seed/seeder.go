package seed

import (
	"log"

	"choosr-server/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "Mike",
		Email:    "mike@gmail.com",
		Password: "password1",
	},
	models.User{
		Username: "Luces",
		Email:    "luces@gmail.com",
		Password: "password2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
