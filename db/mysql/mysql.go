package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// 查询多条
func QueryList() {
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

// 查询一条
func QueryRow() {
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
	row := db.QueryRow("select id, username, age, address from users where username = ?", "张三")
	err = row.Scan(&id, &username, &age, &address)
	if err != nil {
		fmt.Println("query row failed", err)
		return
	}
	fmt.Println(id, username, age, address)
}

// 插入
func InsertOne() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("open conn failed", err)
		return
	}
	defer db.Close()
	ret, _ := db.Exec("insert into users(id, username, age, address) " +
		"values('" + uuid.New().String() + "', 'fenger', 25, '北京市海地区')")
	insID, _ := ret.LastInsertId()
	fmt.Println(insID)
}
