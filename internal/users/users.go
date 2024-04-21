package users

import (
	"database/sql"
	"log"

	"github.com/maskedeman/hackernews/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := mysql.Db.Prepare("INSERT INTO Users(Username, Password) VALUES(?,?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIDByUsername(username string) (int, error) {
	statement, err := mysql.Db.Prepare("SELECT id FROM Users WHERE username=?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int

	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return Id, nil
}
