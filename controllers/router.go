package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
)

// サーバーを起動させるためのエンジンを初期化
var Router *gin.Engine

func init() {
  // デフォルトのエンジンを作成
  router := gin.Default()

  userController := NewUserController(database.NewSqlHandler())

  // ルーティングを追加
  router.GET("/user/get/:id", func(c *gin.Context) { userController.Get(c) })

  Router = router
}
