package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	"time"
	"io"
)

func readBody() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uasge: go run http.go url")
		os.Exit(1)
	}
	url := os.Args[1]
	// 判断字符串是否是以指定子串开头
	if ! strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	// http get请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get %s error: %v", url, err)
		os.Exit(1)
	}
	fmt.Print("status_code = ", resp.Status)
	// 读取流中所有数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read body error: %v", err)
		os.Exit(1)
	}
	// 关闭流
	resp.Body.Close()
	fmt.Print("body = ", string(body))
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Get %s fail: %v", url, err)
		return
	}
	// 相当于sendfile
	nbyte, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("read body fail: %v", err)
		return
	}
	second := time.Since(start).Seconds()
	ch <- fmt.Sprintf("get %s %d byte %f second", url, nbyte, second)

}

func fetchall() {
	start := time.Now()
	// 定义一个chan
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total: %f second\n", time.Since(start).Seconds())
}

func main()  {
//	readBody()
	fetchall()
}