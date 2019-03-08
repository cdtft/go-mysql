package original

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int		//id
	Name   string	//名称
	Age    int		//年龄
	Gender int		//性别0-女 1-男
}

func InitDB() *sql.DB {
	DB, _ := sql.Open("mysql", "root:root@/go-learn?charset=utf8")
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail", err.Error())
		return nil
	}
	fmt.Println("connect success")
	return DB
}

func InsertUser(user *User, db *sql.DB) int64 {
	stmt, err := db.Prepare("insert into `user` (`name`, `age`, `gender`) values (?, ?, ?)")
	if err != nil {
		fmt.Println("prepare fail")
		return -1
	}
	result, _ := stmt.Exec(user.Name, user.Age, user.Gender)
	id, _ := result.LastInsertId()
	return id
}

func SelectAllUser(db *sql.DB) []User {
	stmt, err := db.Prepare("select * from `user`")
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, _ := stmt.Query()
	var userList []User
	for rows.Next() {
		var user User
		//顺序都要对！！！
		_ = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
		userList = append(userList, user)
	}
	return userList
}
