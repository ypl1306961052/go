package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func f1(w http.ResponseWriter, r *http.Request) {
	by, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println("打开文件失败")
		w.Write([]byte("打开文件失败"))
		return
	}
	count, err := w.Write(by)
	if err != nil {
		fmt.Println("相应失败")
		return
	}
	fmt.Println("相应的字节数量:", count)

}

func add(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"]
	fmt.Println(id)
	by, _ := ioutil.ReadFile("./index.html")
	byStr := string(by)
	byStr = strings.Replace(byStr, "#{id}", id[0], 1)
	count, _ := w.Write([]byte(byStr))
	fmt.Println("相应的字节数量:", count)
}

func UnloadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("保存图片....")
	imageName := r.Header.Get("imageName")
	fmt.Println(imageName)
	if imageName == "" {
		imageName = "image.png"
	}
	imageBy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("./"+imageName, imageBy, os.ModePerm)
	if err != nil {
		fmt.Println("保存文件失败")
	}

	_, err = w.Write([]byte("保存文件成功"))
	if err != nil {
		fmt.Println("相应客户端失败,", err)
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			fmt.Println("关闭请求失败,", err)
		}
		fmt.Println("保存图片完毕")

	}()
}
func form(w http.ResponseWriter, r *http.Request) {
	by, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("head:", r.Header)
	va, _ := url.ParseQuery(string(by))
	fmt.Println(va)
	r.ParseForm()
	fmt.Println("body:", r.Form)
	println(r.Form.Get("name"))
	fmt.Println("name:", r.PostFormValue("name"))
	fmt.Println("name:", r.FormValue("name"))

	count, err := w.Write([]byte("收到form"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("相应数量:", count)
	defer r.Body.Close()
}

func main() {
	service1()
}
func service1() {

	fmt.Println("/add")
	http.HandleFunc("/add", add)
	http.HandleFunc("/index", f1)
	fmt.Println("/index")
	http.HandleFunc("/upload", UnloadImage)
	fmt.Println("/upload")
	http.HandleFunc("/from", form)
	fmt.Println("/from")
	fmt.Println("端口:9090 已启动")
	log.Fatal(http.ListenAndServe("127.0.0.1:9090", nil))
}

//自定义服务参数
func service2() {
	ser := http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	ser.ListenAndServe()

}
