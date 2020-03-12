package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Person struct {
	age   int
	name  string
	adder string
}

var wg sync.WaitGroup

func main() {
	initDB()
	defer closeDb()
	////goDb()
	////selectColumns()
	////Transaction()
	createTablePerson()
	////showTables()
	initDataToPerson()
	////selectData()
	////print(*createLineValue("abv", "23"))
	//createTableUserID()
	//insertDataToUserId()
	//u1, _ := uuid.NewV4()
	//println(u1.String())
	//randName()
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

func createTablePerson() {
	_, er := db.Exec("create table person(\n" +
		"id int not null auto_increment,\n" +
		"first_name varchar (32) ,\n" +
		"last_name varchar (32) ,\n" +
		"birthday date ,\n" +
		"city varchar (50),\n" +
		"key firstName(`first_name`),\n" +
		"primary key (id)\n" +
		");")
	if er != nil {
		fmt.Println("crate table person error", er)
		panic("create table person error")
	} else {
		fmt.Println("create table success")
	}

}
func showTables() {
	re, er := db.Exec("show tables")
	if er != nil {
		fmt.Println("crate table person error", er)
	} else {

		fmt.Println(re.RowsAffected())
	}
}

func initDataToPerson() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go initInnerDataToPerson(10000, i)
	}
	wg.Wait()
}
func initInnerDataToPerson(count int, index int) {
	char := initCharacter()
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		firstRand := rand.Intn(3)
		rand.Seed(time.Now().UnixNano())
		lastRand := rand.Intn(3)
		_, er := db.Exec("insert into person(first_name,last_name,birthday,city)"+
			"value(?,?,?,?)", naming(3+firstRand, char), naming(3+lastRand, char), randBirthday(), randCity())
		if er != nil {
			fmt.Println("insert into error,", er)
		}
	}
	fmt.Println(index, "执行完毕")
	wg.Done()

}
func randBirthday() string {

	rand.Seed(time.Now().UnixNano())
	year := 1970 + rand.Intn(40)
	month := time.Month(rand.Intn(13))
	rand.Seed(time.Now().UnixNano())
	day := rand.Intn(29)
	timeStr := time.Date(year, month, day, 0, 0, 0, 0, time.Local, )
	return timeStr.Format("2006-01-02")

}
func randCity() string {
	citys := [...]string{"上海", "北京", "成都", "海口"}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(citys))
	return citys[index]
}
func naming(count int, char *[26]string) string {
	var name = ""
	//加盐

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(26)
		name = name + char[index]
	}
	return name
}
func randName() {
	char := initCharacter()
	names := make([]string, 0)
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		name := naming(3+rand.Intn(3), char)
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println(names)

}

func initCharacter() *[26]string {
	character := [26]string{}
	bytes := byte('a')
	character[0] = string(bytes)
	for i := 1; i < 26; i++ {
		bytes = bytes + 1
		character[i] = string(bytes)
	}
	return &character
}
func selectData() {
	//start := time.Now().UnixNano()
	rows, er := db.Query("select first_name,last_name,birthday,city from person where first_name=?;", "anvhd")
	defer rows.Close()
	if er != nil {
		fmt.Println("select data error")
	}
	var firstName = ""
	var lastName = ""
	var birthday = ""
	var city = ""
	if rows == nil {
		return
	}
	for rows.Next() {

		err := rows.Scan(&firstName, &lastName, &birthday, &city)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(firstName, lastName, birthday, city)

	}
	//fmt.Println((time.Now().UnixNano()-start)/1000000,"ms")
}
func createTableUserID() {
	_, err := db.Exec("create table userId(\n" +
		"id int not null auto_increment,\n" +
		"name varchar (32)  not null default '',\n" +
		"age tinyint(1) unsigned not null default 0,\n" +
		"primary key(id),\n" +
		"key firstName(name)\n" +
		");")
	if err != nil {
		fmt.Println("create table userId error,", err)
	}
}

func createTableUserUUid() {
	_, err := db.Exec("create table userUUid(\n" +
		"uuid varchar(100) not null ,\n" +
		"name varchar (32)  not null default '',\n" +
		"age tinyint(1) unsigned not null default 0,\n" +
		"primary key(uuid),\n" +
		"key firstName(name)\n" +
		");")
	if err != nil {
		fmt.Println("create table userUUid error,", err)
	}
}
func insertDataToUserId() {
	start := time.Now().Unix()
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go innerInsertDataToUserID(1000)
	}
	wg.Wait()
	fmt.Println("id:", time.Now().Unix()-start, "s")
}

func innerInsertDataToUserID(size int) {
	defer wg.Done()
	sqlInsert := "insert into userId(name,age)" + " values "
	ch := initCharacter()
	var lineValue = ""
	for i := 0; i < size; i++ {

		name := naming(10, ch)
		rand.Seed(time.Now().UnixNano())
		age := strconv.Itoa(rand.Intn(100))
		line := createLineValue(name, age)
		if i == size-1 {
			str := "('" + name + "','" + age + "');"
			line = &str
		}
		lineValue = lineValue + *line
	}
	if lineValue == "" {
		fmt.Println("插入语句有误")
		return
	}
	sqlInsert = sqlInsert + lineValue
	re, err := db.Exec(sqlInsert)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(re.RowsAffected())
	}

}

func createLineValue(name string, age string) *string {
	str := "('" + name + "','" + age + "'),"
	return &str
}
func insertDataToUserUUid() {
	start := time.Now().UnixNano()
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go innerInsertDataToUserUUid(1000)
	}
	wg.Wait()
	fmt.Println("uuid:", time.Now().UnixNano()-start/1000000, "ms")
}
func innerInsertDataToUserUUid(size int) {
	defer wg.Done()
	sqlInsert := "insert into userUUid(uuid,name,age)" + " values "
	ch := initCharacter()
	var lineValue = ""
	for i := 0; i < size; i++ {
		u1, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
			continue
		}
		name := naming(10, ch)
		rand.Seed(time.Now().UnixNano())
		age := strconv.Itoa(rand.Intn(100))
		line := createLineValueUUId(u1.String(), name, string(age))
		if i == size-1 {
			str := "('" + u1.String() + "','" + name + "','" + string(age) + "');"
			line = &str
		}
		lineValue = lineValue + *line
	}
	if lineValue == "" {
		fmt.Println("插入语句有误")
		return
	}
	sqlInsert = sqlInsert + lineValue
	re, err := db.Exec(sqlInsert)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(re.RowsAffected())
	}
}

func createLineValueUUId(uuid string, name string, age string) *string {
	str := "('" + uuid + "','" + name + "','" + age + "'),"
	return &str
}
func mainUseUUid() {
	createTableUserUUid()
	insertDataToUserUUid()
}
