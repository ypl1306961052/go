package main

import (
	"fmt"
	"testing"
)

func TestGoDb(t *testing.T) {
	t.Log("测试成功")

}
func TestF1(t *testing.T) {
	f1()
}
func TestSelectOne(t *testing.T) {
	selectOne(34)

}

//基准测试
func BenchmarkSelectOne(b *testing.B) {
	//initDB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		selectOne(34)
	}
	b.Log("测试结束")

}

//并行测试
func BenchmarkSelectOneParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectOne(34)
		}
	})
}

//1132           1113969 ns/op             497 B/op         13 allocs/op
func BenchmarkInsetDb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insetDb()
	}
	b.Log("测试插入数据结束")
}

//4087            727315 ns/op             499 B/op         13 allocs/op
func BenchmarkInsetDbParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insetDb()
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	initDB()
	fmt.Println("初始化数据库链接")
	m.Run()
	fmt.Println("测试收尾工作开始")
	//clearData()
	closeDb()

}
func BenchmarkSelectData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectData()
	}
}

func clearData() {
	rest, err := db.Exec("delete from Person where age=1")
	checkErr("清楚数据错误", err)
	count, err := rest.RowsAffected()
	checkErr("", err)
	println(count)

}
func TestInsertUserId(t *testing.T) {
	createTableUserID()
	insertDataToUserId()
}

func TestInsertUserUUid(t *testing.T) {
	mainUseUUid()
}