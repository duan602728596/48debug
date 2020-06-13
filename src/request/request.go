package FortyEight_Debug

import (
  "io"
  "io/ioutil"
  "net/http"
)

type RequestOptions struct {
  Url string
  Method string
  Headers map[string]string
  Body io.ReadCloser
}

/**
 * 请求数据
 * @param { RequestArgument } options: 请求配置
 */
func Request(options RequestOptions) []byte {
  client := &http.Client {}
  request, _ := http.NewRequest(options.Method, options.Url, options.Body)

  for key := range options.Headers {
    request.Header.Add(key, options.Headers[key])
  }

  response, _ := client.Do(request)

  defer response.Body.Close()

  resBody, _ := ioutil.ReadAll(response.Body)

  return resBody
}