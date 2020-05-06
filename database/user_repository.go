package database

import "techtrain-CA/models"

// クエリを実行するための構造体
type UserRepository struct {
  SqlHandler *SqlHandler
}

// クエリを実行して結果を返す
func (repo *UserRepository) FindByToken(identifier int) (user models.User, err error) {
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
