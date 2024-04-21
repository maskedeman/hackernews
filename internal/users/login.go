package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maskedeman/hackernews/internal/pkg/db/migrations/mysql"
)

func (user *User) Login() bool {
	stmt, err := mysql.Db.Prepare("select Password from Users where Username=?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(user.Username)

	var hashedPassword string

	err = row.Scan(&hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return CompareHash(user.Password, hashedPassword)
}
