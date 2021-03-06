/*
Package wechat provide wechat sdk for golang

5行代码，链式消息，快速开启微信API示例:

	package main

	import (
		"net/http"
		"github.com/esap/wechat" // 微信SDK包
	)

	func main() {
		wechat.Debug = true
		wechat.Set("esap", "yourAppID", "yourSecret", "yourAesKey")
		http.HandleFunc("/", WxHandler)
		http.ListenAndServe(":9090", nil)
	}

	func WxHandler(w http.ResponseWriter, r *http.Request) {
		wechat.VerifyURL(w, r).NewText("正在查询...").Reply().NewText("客服消息1").Send().NewText("客服消息2").Send()
	}

More info:https://github.com/esap/wechat

*/
package wechat
