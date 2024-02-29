// package main

// import (
// 	"time"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      uint
// 	Birthday time.Time
// }

// func main() {
// 	// Create a new SQLite database file named "mydatabase.db"
// 	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// 	db.Create(&user) // pass pointer of data to Create

// 	// user.ID             // returns inserted data's primary key
// 	// result.Error        // returns error
// 	// result.RowsAffected // returns inserted records count

// 	users := []*User{
// 		&User{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
// 		&User{Name: "Jackson", Age: 19, Birthday: time.Now()},
// 	}

// 	db.Create(users) // pass a slice to insert multiple row

//		// result.Error        // returns error
//		// result.RowsAffected // returns ins
//	}
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_ID  uint
	Name     string
	Age      uint
	Birthday time.Time
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
		&User{User_ID: 1, Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		&User{User_ID: 2, Name: "Jackson", Age: 19, Birthday: time.Now()},
	}

	users = []*User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, users := range users {
		fmt.Println(users) // 1,2,3
	}
	fmt.Println(db.First(&users))
	fmt.Println(db.Last(&user))
	fmt.Println(db.Find(&users, []int{1, 2, 3}))
	db.Find(&users)
	db.Delete(&users, []int{4, 5})

}
