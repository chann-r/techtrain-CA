package main

import (
  // モジュール名/独自パッケージ名でimport
  "techtrain-CA/controllers"
)

func main() {
  controllers.Write("a")
  controllers.Router.Run()
}
