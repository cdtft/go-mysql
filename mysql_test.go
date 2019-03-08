package mysqlDemo

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	_ = InitDB()
}

func TestInsertUser(t *testing.T) {
	db := InitDB()
	user := new(User)
	user.Gender = 1
	user.Name = "nihao"
	user.Age = 23
	InsertUser(user, db)
}

func TestSelectAllUser(t *testing.T) {
	users := SelectAllUser(InitDB())
	for _, user := range users {
		fmt.Printf("%+v", user)
	}
}