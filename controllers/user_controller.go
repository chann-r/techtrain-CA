package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
  "techtrain-CA/models"
  "strconv"
)

// ルーティング処理するための構造体
type UserController struct{
  UserRepository database.UserRepository
}

// データベースと接続するための関数
func NewUserController(sqlHandler *database.SqlHandler) *UserController {
  return &UserController{
    UserRepository: database.UserRepository{
      SqlHandler: sqlHandler,
    },
  }
}

// POSTリクエストに対して、ユーザーを保存して保存したユーザーのトークンをJSONで返す
func (controller *UserController) Create(c *gin.Context) {
  u := models.User{}

  // リクエストの Content-Typeをチェックしてバインド（JsonとXML以外だとエラーを吐く）
  err := c.Bind(&u)

  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // 保存して、保存したidを取得
  id, err := controller.UserRepository.Store(u)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // idを元に User を検索
  user, err := controller.UserRepository.FindById(id)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // トークンを作成
  tokenString, err := controller.UserRepository.CreateToken(user)

  if err != nil {
		c.JSON(500, err.Error())
		return
	}
  // トークンを返す
	c.JSON(200, tokenString)
}

// GETリクエストがきたら、クエリからパラメーターを取得して、処理してJSONで返す
func (controller *UserController) GetUser(c *gin.Context) {
  // Paramメソッドでクエリのidを取得し、Atoiメソッドでintに変換
  id, err := strconv.Atoi(c.Param("id"))

  if err != nil {
    c.JSON(500, err)
  }

  user, err := controller.UserRepository.FindById(id)
  if err != nil {
    c.JSON(500, err)
    return
  }
  c.JSON(200, user)
}
