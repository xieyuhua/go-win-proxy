package main

import (
	"flag"
	"fmt"
	"sysproxy/gosysproxy"
	"time"
)

/*
Golang 配置Windows系统代理。

// SetPAC 设置PAC代理模式
// scriptLoc: 脚本地址，如: "http://127.0.0.1:7777/pac"
func SetPAC(scriptLoc string)

// SetGlobalProxy 设置全局代理
// proxyServer: 代理服务器host:port，例如: "127.0.0.1:7890"
// bypass: 忽略代理列表,这些配置项开头的地址不进行代理
func SetGlobalProxy(proxyServer string, bypasses ...string) error

// Off 关闭代理
func Off() error

// Flush 更新系统配置使生效
func Flush()

// Status 获取当前系统代理配置
func Status() (*ProxyStatus, error)
*/
func main() {
	ip := flag.String("ip", "127.0.0.1:1080", "proxy 127.0.0.1:1080")
	flag.Parse()
	// 设置全局代理  SetGlobalProxy("127.0.0.1:7890", "foo", "bar", "<local>")
	err := gosysproxy.SetGlobalProxy(*ip)
	if err != nil {
		panic(err)
	}

	res, errs := gosysproxy.Status()
	if errs != nil {
		panic(errs)
	}
	switch res.Type {
	case gosysproxy.INTERNET_OPEN_TYPE_PRECONFIG:
		fmt.Println(">> 代理方式：", "use registry configuration")
	case gosysproxy.INTERNET_OPEN_TYPE_DIRECT:
		fmt.Println(">> 代理方式：", "不代理 direct to net")
	case gosysproxy.INTERNET_OPEN_TYPE_PROXY:
		fmt.Println(">> 代理方式：", "使用代理服务器")
	}
	fmt.Println(">> 代理地址：", res.Proxy)
	fmt.Println(">> 请勿对以下列条目开头的地址使用代理服务器：", res.Bypass)
	fmt.Println(">> 请勿将代理服务器用于本地(Intranet)地址：", res.DisableProxyIntranet)

	time.Sleep(time.Second * 60)

	err = gosysproxy.Off()
	if err != nil {
		panic(err)
	}
}
