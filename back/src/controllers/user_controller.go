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

  // トークンを格納したUser
  user, err := controller.UserRepository.CreateToken(u)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}

  // 保存して、保存したidを取得
  _, err = controller.UserRepository.Store(user)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  // tokenのmapを作成
  token := map[string]string{"token": user.Token}

  // トークンを返す
	c.JSON(200, token)
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

// ヘッダーのtokenを取得してデータベースと照合して、ユーザー名をJSONで返す
func (controller *UserController) Get(c *gin.Context) {
  // ヘッダーのtokenを取得
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

  // nameのmapを作成
  name := map[string]string{"name": user.Name}

  c.JSON(200, name)
}

// ヘッダーのtokenを取得してリクエストボディにNameに従ってユーザー名を変更する
func (controller *UserController) Update(c *gin.Context) {
  u := models.User{}

  err := c.Bind(&u)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  tokenString := c.Request.Header.Get("x-token")
  if tokenString == "" {
    c.JSON(500, "token must be needed.")
    return
  }

  user, err := controller.UserRepository.FindByToken(tokenString)

  // 変更するユーザーのidと変更情報を渡してユーザー情報を更新
  err = controller.UserRepository.Change(user.Id, u)
  if err != nil {
    c.JSON(500, err.Error())
    return
  }

  c.Status(200)
}
