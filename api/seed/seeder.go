package seed

import (
	// "log"

	"github.com/jinzhu/gorm"
	"apiservice/api/models"
)

var users = []models.User{
	models.User{
		Name: "Seed User one",
		Email:    "test@gmail.com",
		Password: "password",
		Age: 12,
		FavoriteColor: "Red",
		FavoriteOperatingSystem: "Mac OS",
	},
	models.User{
		Name: "Seed User two",
		Email:    "test@email.com",
		Password: "passwordOne",
		Age: 23,
		FavoriteColor: "blue",
		FavoriteOperatingSystem: "Linux",
	},
}

func Load(db *gorm.DB) {

	// err := db.Debug().DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }

	// for i := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}
	// }
}