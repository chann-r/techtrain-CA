package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
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

// GETリクエストがきたら、クエリからパラメーターを取得して、処理してJSONで返す
func (controller *UserController) Get(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  user, err := controller.UserRepository.FindByToken(id)
  if err != nil {
    c.JSON(500, err)
    return
  }
  c.JSON(200, user)
}
