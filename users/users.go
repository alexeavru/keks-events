package users

import (
	"database/sql"
	"log"

	"github.com/alexeavru/keks-events/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func CheckPassword(username string, password string) (bool, error) {

	row, err := database.Db.Query("SELECT password FROM users WHERE name = $1", username)
	if err != nil {
		log.Fatal(err)
	}

	var hashedPassword string
	for row.Next() {
		err = row.Scan(&hashedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return false, nil
			} else {
				log.Fatal(err)
			}
		}
	}

	if CheckPasswordHash(password, hashedPassword) {
		log.Printf("Check user password success. User: %s Token generated", username)
		return true, nil
	}
	log.Printf("Error check user password. User and Password not found!")

	return false, err
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
