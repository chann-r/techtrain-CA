package database

import (
  "github.com/dgrijalva/jwt-go"
  "techtrain-CA/models"
  )

// クエリを実行するための構造体
type UserRepository struct {
  SqlHandler *SqlHandler
}

// 保存して保存した行のidを返す
func (repo *UserRepository) Store(user models.User) (id int, err error) {
  result, err := repo.SqlHandler.Execute("INSERT INTO users (name) VALUES (?)", user.Name,)

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
  row, err := repo.SqlHandler.Query("SELECT id, name, created_at, updated_at FROM users WHERE id = ?", identifier)

  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var name string
  var created_at string
  var updated_at string

  row.Next()
  if err = row.Scan(&id, &name, &created_at, &updated_at); err != nil {
    return
  }

  user.Id = id
  user.Name = name
  user.CreatedAt = created_at
  user.UpdatedAt = updated_at

  return
}

// 署名するときに使う鍵
var KEY []byte = []byte("key")

// 署名して生成したトークンを返す
func (repo *UserRepository) CreateToken(user models.User) (token string, err error) {
  jwtToken := models.JwtToken{}

  // ペイロードを作成
  claims := jwt.StandardClaims{
    // claim を設定
    Issuer: "__init__",
  }

  // 署名前の（Header, Claims, Method）が入ったToken
  jwtToken.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // 署名してトークンを生成
  token, err = jwtToken.SignedString(KEY)
  return
}
