package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
  "techtrain-CA/models"
)

type CharacterController struct {
  CollectionRepository database.CollectionRepository
  UserRepository       database.UserRepository
}

func NewCharacterController(sqlHandler *database.SqlHandler) *CharacterController {
  return &CharacterController{
    CollectionRepository: database.CollectionRepository{
      SqlHandler: sqlHandler,
    },
    UserRepository: database.UserRepository{
      SqlHandler: sqlHandler,
    },
  }
}

// トークンでユーザーを識別して、そのユーザーが持っているキャラクター一覧を返す
func (controller *CharacterController) List(c *gin.Context) {
  tokenString := c.Request.Header.Get("x-token")
  if tokenString == "" {
    c.JSON(500, "token must be needed.")
    return
  }

  user, err := controller.UserRepository.FindByToken(tokenString)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  // ユーザーidで所有キャラクターを検索
  userCharacters, err := controller.CollectionRepository.FindByUserId(user.Id)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  // マップに保存したガチャ内容を格納
  characters := map[string]models.UserCharacters{"characters": userCharacters}

  c.JSON(200, characters)
}
