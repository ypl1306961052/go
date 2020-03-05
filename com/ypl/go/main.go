package main

import (
	"com/ypl/go/model"
	json2 "encoding/json"
	"fmt"
	"sync"
	"time"
)

//func main() {
//	var p=&model.Person{
//		Man:model.Man{Name:"杨沛霖"},
//		Age:23,
//		Adder:model.Adder{
//			City:     "上海",
//			Province: "上海",
//		},
//
//	}
//	printAdder(p)
//	peronJson:=json(p)
//	fmt.Println(peronJson)
//	p.SayName()
//
//}
func printAdder(m model.PrintAdder) {
	m.PrintAdder()
}
func json(person *model.Person) string {
	b, error := json2.Marshal(person)
	if error != nil {
		fmt.Println("json解析失败,失败的原因", error)
	}
	return string(b)
}

var wg sync.WaitGroup
var ch chan int

func getChannel(ch *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("等待中")
	time.Sleep(time.Second)
	var index = 0
	//　一直等待　管道中有值
	for {
		if index == 10 {
			fmt.Println("获取数据完毕")
			break
		}
		index = index + 1
		x, ok := <-*ch

		if !ok {
			fmt.Println("获取数据失败")
			break
		} else {
			fmt.Println("从管道的获取的数据为:", x)
		}

	}

}
func noChannelBf() {
	ch = make(chan int, 10)
	//关闭通断
	defer close(ch)
	wg.Add(1)
	go getChannel(&ch, &wg)
	fmt.Println("赋值了")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//fmt.Println(ch)
	fmt.Println("程序结束")
	wg.Wait()

}
func initData(ch *chan int) {
	defer wg.Done()
	fmt.Println("初始化管道")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		*ch <- i
	}
	close(*ch)
	fmt.Println("初始化管道完毕")
}

//ch1 只能读取通道　　　ch2 只能写值
func intoCh2(ch1 *<-chan int, ch2 *chan<- int) {
	defer wg.Done()
	for {
		tmp, ok := <-*ch1
		if !ok {
			fmt.Println("ch1读取数据完毕")
			break
		}
		//tmp = tmp * tmp
		fmt.Println("ch1:", tmp)
		*ch2 <- tmp
	}
	once.Do(func() {
		close(*ch2)
	})
}

var once sync.Once

func getCh2(ch2 *chan int, title string) {
	defer wg.Done()
	for {
		//对于关闭的管道　任然可以读取值,不能写数据了　但是ｏｋ 为false
		tmp, ok := <-*ch2
		if !ok {
			fmt.Println("ch2读取数据完毕")
			break
		} else {
			fmt.Println(title, "_ch2:", tmp)
		}

	}

}
func main() {
	wg.Add(3)
	var ch1 = make(chan int, 100)
	var ch2 = make(chan int, 100)
	go initData(&ch1)
	go intoCh2(&ch1, &ch2)
	go getCh2(&ch2, "1")
	go getCh2(&ch2, "2")
	fmt.Println("执行完毕")
	wg.Wait()
	//noChannelBf()

}

//func main() {
//	//gmp
//
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go thread.SayHello(i, &wg)
//	}
//	fmt.Println("main")
//	wg.Wait()
//	fmt.Println("执行完毕")
//	var p = new(model.Person)
//	fmt.Println(p)
//
//	i := make([]int, 10)
//	m := make(map[string]interface{}, 10)
//	fmt.Println(i)
//	fmt.Println(m)
//	for i := 0; i <= 10; i++ {
//		m[string(i)] = string(i)
//	}
//	fmt.Println(m)
//
//	ch := make(chan int, 10)
//	fmt.Println(ch)
//
//}
