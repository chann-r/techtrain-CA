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

// トークンからuser_idを取得してランダムにcharacter_idを生成して保存して返す
func (controller *GachaController) Draw(c *gin.Context) {
  // リクエストに合う構造体を定義
  type GachaTimes struct {
    Times int
  }
  gachaTimes := GachaTimes{}

  err := c.Bind(&gachaTimes)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

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

  // ユーザーidとランダムで生成したキャラクターidを保存して、collectionのidを返す
  characterIds, err := controller.CollectionRepository.Store(user.Id, gachaTimes.Times)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  // 保存したcollectionを返す
  // collection, err := controller.CollectionRepository.FindById(id)

  c.JSON(200, characterIds)
}
