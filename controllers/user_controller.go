package controllers

import "techtrain-CA/database"

type UserController struct{
  SqlHandler databases.SqlHandler
}

func NewUserController(sqlHandler databases.SqlHandler) * UserController {
  return &UserController{
    SqlHandler: sqlHandler
  }
}
