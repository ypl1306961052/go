package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	//get()
	//Values()
	//post()
	//postForm()
	//head()
	//customHttp1()
	customHttp2()
}

func Values() {
	apiUrl := "http://127.0.0.1:9090/add"
	data := url.Values{}
	data.Set("id", "1000")

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("解析url失败")
		return
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("请求失败")
		return
	}
	defer resp.Body.Close()
	by, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(by))
}
func post() {
	file, err := os.OpenFile("./a.jpg", os.O_RDONLY, 077)
	if err != nil {
		fmt.Println("获取图片错误,", err)
	}
	fileBuf := bufio.NewReader(file)
	re, err := http.Post("http://127.0.0.1:9090/upload?imageName=a.jpg", "image/jpeg", fileBuf)
	if err != nil {
		fmt.Println("相应失败,", err)
	}
	reStr, err := ioutil.ReadAll(re.Body)
	if err != nil {
		fmt.Println("相应读取失败")
	}

	fmt.Println(string(reStr))
}
func postForm() {

	re, err := http.PostForm("http://127.0.0.1:9090/from", url.Values{
		"name": []string{"杨沛霖"},
		"age":  []string{"12"},
		"city": []string{"上海"},
	})
	if err != nil {
		fmt.Println("保单请求失败,", err)
	}
	boBy, _ := ioutil.ReadAll(re.Body)
	fmt.Println(string(boBy))

}
func call() {
	runtime.Caller(0)
}
func get() {

	u := "http://127.0.0.1:9090/index"
	re, err := http.Get(u)
	if err != nil {
		fmt.Println("请求消失败,", err)
	}
	defer re.Body.Close()
	io.Copy(os.Stdout, re.Body)

}
func head() {
	re, err := http.Head("http://127.0.0.1:9090/from")
	if err != nil {
		fmt.Println(err)
	}
	reBy, _ := ioutil.ReadAll(re.Body)
	fmt.Println(string(reBy))
}
func ioReaderStdIn() {
	red := bufio.NewReader(os.Stdin)
	for {
		by, _, _ := red.ReadLine()
		fmt.Println(string(by))

	}
}

func customHttp() {
	va := url.Values{
		"name": []string{"杨沛霖"},
	}
	fmt.Println(strings.NewReader(va.Encode()))
	httpRes, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:9090/from", strings.NewReader(va.Encode()))

	if err != nil {
		fmt.Println(err)
	}
	var header = http.Header{}
	header.Add("name", "杨沛霖")
	header.Add("session", "15708989110")
	httpRes.Header = header

	//re, err := http.DefaultClient.Do(httpRes)
	client := &http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       1000,
	}
	re, err := client.Do(httpRes)
	if err != nil {
		fmt.Println(err)
	}
	reBy, _ := ioutil.ReadAll(re.Body)
	fmt.Println(string(reBy))

}

func customHttp1() {
	tr := &http.Transport{
		//关闭长连接
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 60,
	}
	rep, err := client.Get("http://127.0.0.1:9090/index")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rep.StatusCode)
	by, _ := ioutil.ReadAll(rep.Body)
	fmt.Println(string(by))
}

//集成了transport借口
type MyTransport struct {
	Transport http.RoundTripper
}

func (t *MyTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("提前处理")
	return t.transport().RoundTrip(req)
}
func (t *MyTransport) client(timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: t,
		Timeout:   timeout,
	}

}
func customHttp2() {
	tr := MyTransport{}
	client := tr.client(time.Second * 10)

	req, err := client.Get("http://127.0.0.1:9090/index")
	if err != nil {
		fmt.Println("请求失败,", err)
		return
	}
	conBy, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(conBy))

}
