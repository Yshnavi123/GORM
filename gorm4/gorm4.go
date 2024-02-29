package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Role string
}

type Result struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Create a new SQLite database file named "mydatabase.db"
	db, err := gorm.Open(sqlite.Open("newdatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate to create the 'users' table
	db.AutoMigrate(&User{})

	// Insert some sample data
	db.Create(&User{Name: "John", Age: 25, Role: "admin"})
	db.Create(&User{Name: "Jane", Age: 30, Role: "user"})

	// SELECT query with parameter
	var resultByID Result
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 2).Scan(&resultByID)
	fmt.Println("Result by ID:", resultByID)

	// SELECT query with parameter
	var resultByName Result
	db.Raw("SELECT id, name, age FROM users WHERE name = ?", "John").Scan(&resultByName)
	fmt.Println("Result by Name:", resultByName)

	// SELECT query with aggregate function
	var totalAge int
	db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&totalAge)
	fmt.Println("Total Age for Admins:", totalAge)

	// UPDATE query with RETURNING clause
	var updatedUsers []User
	db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "UpdatedName", 30).Scan(&updatedUsers)
	fmt.Println("Updated Users:", updatedUsers)
}
