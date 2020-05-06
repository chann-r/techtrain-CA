package main

import (
  "fmt"
  "rsc.io/quote"
  // モジュール名/独自パッケージ名でimport
  "techtrain-CA/controllers"
)

func main() {
  fmt.Println(quote.Hello())
  controllers.Write("test")
}
