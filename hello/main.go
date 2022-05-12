package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {
	// 数据库操作
	// 逻辑处理
	// restAPI json/xml 返回
	// 1.获取前端传递的参数
	// mobile, passwd
	// 解析参数
	// 如何获得参数
	// 解析参数
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if mobile == "18612345678" && passwd == "123456" {
		loginok = true
	}
	if loginok {
		// {"id":1, "token":"xx"}
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "")
	} else {
		Resp(writer, -1, nil, "密码错误")
	}
	// 返回JSON ok

	// 如何返回JSON

	// io.WriteString(writer, "hello world!")
}

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	// 设置200状态
	w.WriteHeader(http.StatusOK)
	// 输出
	// 定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// 将结构体转化成JSON字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	// 输出
	w.Write(ret)
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)
	// start web server
	http.ListenAndServe("localhost:8080", nil)
}
