package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Query() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("open conn failed", err)
		return
	}
	defer db.Close()

	var id string
	var username string
	var age int
	var address string
	rows, err := db.Query("select id, username, age, address from users")
	if err != nil {
		fmt.Println("query failed", err)
		return
	}
	for rows.Next() {
		rows.Scan(&id, &username, &age, &address)
		fmt.Println(id, username, age, address)
	}
	defer rows.Close()
}
