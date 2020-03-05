package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
)

type Person struct {
	age   int
	name  string
	adder string
}

func main() {
	initDB()
	defer closeDb()
	//goDb()
	//selectColumns()
	Transaction()
}
func f1() {
	fmt.Println("测试")
}
func goDb() {

	selectDb()
}

var db *sql.DB
var once sync.Once

func initDB() {

	once.Do(func() {
		fmt.Println("开启数据库链接.....")
		var err error
		db, err = sql.Open("mysql", "root:ypl123456@tcp(127.0.0.1:3306)/go?charset=utf8")
		if err != nil {
			fmt.Println("数据链接错误", err)
			panic("数据链接错误")
		}
		err = db.Ping()
		if err != nil {
			fmt.Println("数据链接失败,请检查网络")
			print("数据链接失败,请检查网络")
		}
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		fmt.Println("数据库初始化完成")
	})
}
func closeDb() {
	if db != nil {
		err := db.Close()
		if err != nil {

			fmt.Println("关闭数据数据库失败", err)
		}
	}
	fmt.Println("关闭数据库链接")
}
func selectDb() {

	rows, err := db.Query("select * from Person where age>?;", 23)
	defer rows.Close()
	if err != nil {
		fmt.Println("查询数据失败")
	}
	if rows == nil {
		fmt.Println("获取数据错误")
		return
	}
	peoples := make([]Person, 0)
	for rows.Next() {
		var p = Person{}
		err := rows.Scan(&p.age, &p.name, &p.adder)
		if err != nil {
			fmt.Println("获取数据错误")
			continue
		}
		peoples = append(peoples, p)

	}
	fmt.Println(peoples)

}
func selectOne(age int) {

	row := db.QueryRow("select age,name,adder from Person where age=?", age)
	var p = &Person{}
	err := row.Scan(&p.age, &p.name, &p.adder)
	if err != nil {
		fmt.Println("查询失败", err)
	}
	//fmt.Println(p)

}
func insetDb() {

	rest, err := db.Exec("insert into Person(age,name,adder) values (?,?,?)", 1, "杨沛霖", "上海")
	checkErr("", err)
	//id, err := rest.LastInsertId()
	//checkErr("error id,", &err)
	//fmt.Println(id)
	count, _ := rest.RowsAffected()
	fmt.Println("影响行数:", count)
}
func checkErr(lineHead string, err error) {
	if err != nil {
		fmt.Println(lineHead, err)
	}
}
func selectColumns() {
	stmt, err := db.Prepare("select * from Person") //[id uname age mobile]
	checkErr("", err)
	defer stmt.Close()
	rows, err := stmt.Query()
	defer rows.Close()
	checkErr("", err)
	clou, err := rows.Columns()
	clouType, err := rows.ColumnTypes()
	checkErr("", err)
	for _, item := range clouType {
		fmt.Println(item.Name())
		fmt.Println(item.DatabaseTypeName())
		fmt.Println(item.DecimalSize())
		fmt.Println(item.Length())
		fmt.Println(item.Nullable())
	}
	checkErr("", err)
	fmt.Println(strings.Join(clou, ","))

}
func Transaction() {
	t, err := db.Begin()
	checkErr("", err)
	_, err = t.Exec("update  Person set age=age+10 where name='马飞飞'")
	if err != nil {
		fmt.Println("出现错误")
		t.Rollback()
		fmt.Println("事务:1回滚")
		return
	}
	_, err = t.Exec("update  Person set age=age+10 where name='杨kk'")
	if err != nil {
		fmt.Println("出现错误")
		t.Rollback()
		fmt.Println("事务:2回滚")
		return
	}
	t.Commit()
}
