package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "techtrain-CA/database"
  "time"
)

// サーバーを起動させるためのエンジンを初期化
var Router *gin.Engine

// 各種パッケージの init関数は main関数よりも先に呼ばれる
func init() {
  // デフォルトのエンジンを作成
  router := gin.Default()

  router.Use(cors.New(cors.Config{
    // 許可するメソッド
    AllowMethods: []string{
      "POST",
      "GET",
      "PUT",
      "DELETE",
    },
    // 許可するリクエストヘッダー
    AllowHeaders: []string{
      "Access-Control-Allow-Headers",
      "Content-Type",
      "Content-Length",
      "Accept-Encoding",
      "X-CSRF-Token",
      "Authorization",
    },
    // 許可するアクセス元
    AllowOrigins: []string{
      "*",
    },
    MaxAge: 24 * time.Hour,
  }))

  // DBに接続 & コントローラーを初期化
  userController := NewUserController(database.NewSqlHandler())
  gachaController := NewGachaController(database.NewSqlHandler())
  characterController := NewCharacterController(database.NewSqlHandler())

  // ユーザー関連のエンドポイント
  router.POST("/user/create", func(c *gin.Context) { userController.Create(c) })
  router.GET("/user/get/:id", func(c *gin.Context) { userController.GetUser(c) })
  router.GET("/user/get", func(c *gin.Context) { userController.Get(c) })
  router.PUT("/user/update", func(c *gin.Context) { userController.Update(c) })

  // ガチャ関連のエンドポイント
  router.POST("/gacha/draw", func(c *gin.Context) { gachaController.Draw(c) })

  // キャラクター関連のエンドポイント
  router.GET("/character/list", func(c *gin.Context) { characterController.List(c) })

  Router = router
}
