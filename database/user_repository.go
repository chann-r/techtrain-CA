package database

import (
  "github.com/dgrijalva/jwt-go"
  "techtrain-CA/models"
  "time"
  )

// クエリを実行するための構造体
type UserRepository struct {
  SqlHandler *SqlHandler
}

// 保存して保存した行のidを返す
func (repo *UserRepository) Store(user models.User) (id int, err error) {
  result, err := repo.SqlHandler.Execute("INSERT INTO users (name, token) VALUES (?, ?)", user.Name, user.Token)

  if err != nil {
    return
  }

  // 保存した行のidを取得
  identifier, err := result.LastInsertId()
  if err != nil {
    return
  }

  // intに変換
  id = int(identifier)
  return
}

// クエリを実行して結果を返す
func (repo *UserRepository) FindById(identifier int) (user models.User, err error) {
  row, err := repo.SqlHandler.Query("SELECT id, name, token, created_at, updated_at FROM users WHERE id = ?", identifier)

  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var name string
  var token string
  var created_at string
  var updated_at string

  row.Next()
  if err = row.Scan(&id, &name, &token, &created_at, &updated_at); err != nil {
    return
  }

  user.Id = id
  user.Name = name
  user.Token = token
  user.CreatedAt = created_at
  user.UpdatedAt = updated_at

  return
}

// 署名するときに使う鍵
var KEY []byte = []byte("key")

// 署名して生成したトークンを返す
func (repo *UserRepository) CreateToken(u models.User) (user models.User, err error) {
  jwtToken := models.JwtToken{}

  // ペイロードを作成
  claims := jwt.StandardClaims{
    // claim を設定
    Issuer: "__init__",
    Subject: u.Name,
    IssuedAt: time.Now().Unix(),
  }

  // 署名前の（Header, Claims, Method）が入ったToken
  jwtToken.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // 署名してトークンを生成
  tokenString, err := jwtToken.SignedString(KEY)

  if err != nil {
    return
  }

  u.Token = tokenString
  user = u

  return
}

// トークンでユーザーを検索して返す
func (repo *UserRepository) FindByToken(tokenString string) (user models.User, err error) {
  row, err := repo.SqlHandler.Query("SELECT id, name, token FROM users WHERE token = ?", tokenString)

  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var name string
  var token string

  row.Next()
  if err = row.Scan(&id, &name, &token); err != nil {
    return
  }

  user = models.User {
    Id: id,
    Name: name,
    Token: token,
  }

  return
}

// ユーザーidでユーザーを検索して、ユーザー名を更新してエラーを返す
func (repo *UserRepository) Change(id int, user models.User) (err error) {
  row, err := repo.SqlHandler.Query("UPDATE users set name = ? WHERE id = ?", user.Name, id)

  defer row.Close()

  return err
}
