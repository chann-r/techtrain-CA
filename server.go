package main

// モジュール名/独自パッケージ名でimport
import "techtrain-CA/controllers"

func main() {
  controllers.Router.Run()
}
