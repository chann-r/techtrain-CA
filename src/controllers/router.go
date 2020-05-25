package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
)

// サーバーを起動させるためのエンジンを初期化
var Router *gin.Engine

// 各種パッケージの init関数は main関数よりも先に呼ばれる
func init() {
  // デフォルトのエンジンを作成
  router := gin.Default()

  // DBに接続 & コントローラーを初期化
  userController := NewUserController(database.NewSqlHandler())
  gachaController := NewGachaController(database.NewSqlHandler())

  //ユーザー関連のエンドポイント
  router.POST("/user/create", func(c *gin.Context) { userController.Create(c) })
  router.GET("/user/get/:id", func(c *gin.Context) { userController.GetUser(c) })
  router.GET("/user/get", func(c *gin.Context) { userController.Get(c) })
  router.PUT("/user/update", func(c *gin.Context) { userController.Update(c) })

  router.POST("/gacha/draw", func(c *gin.Context) { gachaController.Draw(c) })

  Router = router
}
