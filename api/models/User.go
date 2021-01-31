package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"
	"regexp"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	ID        					uint32    	`gorm:"primary_key;auto_increment"`
	Name  						string    	`gorm:"size:255;not null;"`
	Email     					string    	`gorm:"size:100;not null;unique"`
	Password  					string    	`gorm:"size:100;not null;"`
	Age  						int			`json:"age"`
	FavoriteColor 				string 		`gorm:"size:100;not null;" json:"favorite_color"`
	FavoriteOperatingSystem 	string 		`gorm:"size:100;not null;" json:"favorite_operating_system"`
	CreatedAt 					time.Time 	`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 					time.Time 	`gorm:"default:CURRENT_TIMESTAMP"`
}

// Hash password using bcrypt
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword using bcrypt
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave hook to hash password before save
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare user data before insert
func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FavoriteColor = html.EscapeString(strings.TrimSpace(u.FavoriteColor))
	u.FavoriteOperatingSystem = html.EscapeString(strings.TrimSpace(u.FavoriteOperatingSystem))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate there params before saving
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Name is required")
		}

		if u.Name != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.Name) {
			return errors.New("Name must only contain letters")
		}

		if u.FavoriteColor != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FavoriteColor) {
			return errors.New("FavoriteColor must only contain letters")
		}

		if u.FavoriteOperatingSystem != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FavoriteOperatingSystem) {
			return errors.New("FavoriteOperatingSystem must only contain letters")
		}

		if u.Password == "" {
			return errors.New("Required Password")
		}

		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Email == "" {
			return errors.New("Password is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Name is required")
		}

		if u.Name != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.Name) {
			return errors.New("Name must only contain letters")
		}

		if u.FavoriteColor != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FavoriteColor) {
			return errors.New("FavoriteColor must only contain letters")
		}

		if u.FavoriteOperatingSystem != "" && !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FavoriteOperatingSystem) {
			return errors.New("FavoriteOperatingSystem must only contain letters")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  			u.Password,
			"name":  	 			u.Name,
			"email":     			u.Email,
			"age":       			u.Age,
			"favorite_color":       u.FavoriteColor,
			"favorite_operating_system":     u.FavoriteOperatingSystem,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}