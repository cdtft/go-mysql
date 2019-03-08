###GO语言数据库操作

####original
最原始的连接数据库的操作
```go
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
```
让我想到了当年学Java的时候写的JDBC，在做查询的时候
```go
_ = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
```
Scan中的顺序要和表中字段的顺序相同
