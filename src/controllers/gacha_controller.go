package controllers

import (
  "github.com/gin-gonic/gin"
  "techtrain-CA/database"
)

type GachaController struct {
  CollectionRepository database.CollectionRepository
}

func NewGachaController(sqlHandler *database.SqlHandler) *GachaController {
  return &GachaController{
    CollectionRepository: database.CollectionRepository{
      SqlHandler: sqlHandler,
    },
  }
}

func (controller *GachaController) Draw(c *gin.Context) {
  c.JSON(200, "a")
}
