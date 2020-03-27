package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://172.17.0.7:5001/items"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
