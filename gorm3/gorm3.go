package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_ID  uint
	Name     string
	Age      uint
	Birthday time.Time
	UUID     string // Assuming UUID field
	Role     string // Assuming Role field
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	if u.Role == "admin" {
		return fmt.Errorf("invalid role: %s", u.Role)
	}
	return
}

func main() {
	// Create a new SQLite database file named "mydatabase.db"
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema to create or update the 'users' table
	db.AutoMigrate(&User{})

	user := User{User_ID: 1, Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.Create(&user)

	users := []*User{
		{User_ID: 2, Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{User_ID: 3, Name: "Jackson", Age: 19, Birthday: time.Now()},
	}

	// Create users with create hook
	for _, user := range users {
		db.Create(user)
	}

	for _, user := range users {
		fmt.Println(user)
	}

	db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
}
