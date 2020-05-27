package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
)

type GachaController struct {
  CollectionRepository database.CollectionRepository
  UserRepository       database.UserRepository
}

func NewGachaController(sqlHandler *database.SqlHandler) *GachaController {
  return &GachaController{
    CollectionRepository: database.CollectionRepository{
      SqlHandler: sqlHandler,
    },
    UserRepository: database.UserRepository{
      SqlHandler: sqlHandler,
    },
  }
}

func (controller *GachaController) Draw(c *gin.Context) {
  // ヘッdサーのtokenを取得
  tokenString := c.Request.Header.Get("x-token")
  if tokenString == "" {
    c.JSON(500, "token must be needed.")
    return
  }

  // トークンでユーザーを検索
  user, err := controller.UserRepository.FindByToken(tokenString)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  c.JSON(200, user)
}
