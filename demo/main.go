/*
* 作者：刘时明
* 时间：2020/8/16 0016-11:59
* 作用：
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

var ok int32
var errors int32
var fail int32

func main() {
	all, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(all), "\r\n")
	var wg sync.WaitGroup

	var count int
	for _, v := range arr {
		if strings.HasPrefix(v, "http") {
			go httpHandler(v, &wg)
			count++
		} else {
			go httpHandler("http://"+v, &wg)
			go httpHandler("https://"+v, &wg)
			count += 2
		}
	}
	wg.Wait()

	fmt.Println("统计结果：")
	fmt.Println("请求总数：", count)
	fmt.Println("访问成功：", ok)
	fmt.Println("访问失败：", fail)
	fmt.Println("连接异常：", errors)
}

func httpHandler(url string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " 返回异常，错误为", err.Error())
		atomic.AddInt32(&errors, 1)
		return
	}
	defer response.Body.Close()
	if response.Status == "200 OK" {
		atomic.AddInt32(&ok, 1)
	} else {
		atomic.AddInt32(&fail, 1)
	}
	fmt.Println(url, " 返回code=", response.Status)
}
