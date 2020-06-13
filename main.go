package main

import (
  "encoding/json"
  "fmt"
  "github.com/AlecAivazis/survey/v2"
  request "FortyEight_Debug/src/request"
  "io/ioutil"
  pocket48 "FortyEight_Debug/src/pocket48"
  "strings"
)

func main()  {
  // 账号输入
  username := "" // 用户名（手机号）
  password := "" // 密码

  fmt.Println("请输入口袋48用户名（手机号）和密码")

  survey.AskOne(&survey.Input { Message: "用户名：" }, &username)
  survey.AskOne(&survey.Password{ Message: "密  码：" }, &password)

  if username == "" || password == "" {
    return
  }

  // 请求数据
  body := map[string]string {
    "mobile": username,
    "pwd": password,
  }
  bodyJson, _ := json.Marshal(body)
  bodyString := string(bodyJson)

  // 用户登陆
  var options request.RequestOptions

  options.Url = "https://pocketapi.48.cn/user/api/v1/login/app/mobile"
  options.Method = "POST"
  options.Headers = map[string]string {
    "appInfo": "{\"vendor\":\"apple\",\"deviceId\":\"8OX3L7YQ-0VAD-3I2V-425M-M0S5FBDATD4K\"," +
      "\"appVersion\":\"6.0.1\",\"appBuild\":\"190420\",\"osVersion\":\"11.4.1\",\"osType\":\"ios\",\"" +
      "deviceName\":\"iPhone 6s\",\"os\":\"ios\"}",
    "Content-Type": "application/json; charset=utf8",
  }
  options.Body = ioutil.NopCloser(strings.NewReader(bodyString))

  resBody := request.Request(options)

  // 解析json
  var data pocket48.LoginResponseData

  json.Unmarshal([]byte(resBody), &data)

  // 输出
  fmt.Println("用户ID ", data.Content.UserInfo.UserId)
  fmt.Println("用户名 ", data.Content.UserInfo.Nickname)
}