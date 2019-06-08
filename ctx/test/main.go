package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	//cfg := rice.MustFindBox("conf")
	//f, err := cfg.Open("config-default.ini")
	//d, err := ioutil.ReadAll(f)
	//fmt.Println(string(d), err)
	//a, b := rice.FindBox("config.ini")
	//fmt.Println(a, b)
	//合建chan
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)
	//阻塞直到有信号传入
	fmt.Println("启动")
	go func() {
		s := <-c
		fmt.Println("退出信号", s)
	}()
	time.Sleep(4 * time.Second)
}
