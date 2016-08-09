package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func markup(value interface{}) interface{} {
	if value != nil {
		return value
	}
	return "nil"
}

func main() {
	db, err := sql.Open("mysql", "root:1234@/seezero?charset=utf8")
	fmt.Println(db)
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM users")

	checkErr(err)

	for rows.Next() {
		var id int
		var email string
		var password string
		var username string
		err = rows.Scan(&id, &email, &password, &username)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(email)
		fmt.Println(username)
		fmt.Println(password)
	}

}
