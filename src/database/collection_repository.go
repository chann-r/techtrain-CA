package database

import (
  "math/rand"
  "time"
  "techtrain-CA/models"
)

type CollectionRepository struct {
  SqlHandler *SqlHandler
}

// ランダムにcharacter_idを生成してuser_idと一緒に保存する関数
func (repo *CollectionRepository) Store(user_id int) (id int, err error) {
  // シードを与える（デフォルトだと同じ乱数ジェネレーターを使用してしまう）
  rand.Seed(time.Now().UnixNano())
  // 0から2までの乱数に1を足す
  character_id := rand.Intn(2) + 1

  result, err := repo.SqlHandler.Execute("INSERT INTO collections (user_id, character_id) VALUES (?, ?)", user_id, character_id)
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

// collectionのidでcollectionを検索して返す関数
func (repo *CollectionRepository) FindById(identifier int) (collections models.Collection, err error) {
  row, err := repo.SqlHandler.Query("SELECT id, user_id, character_id FROM collections WHERE id = ?", identifier)

  defer row.Close()

  if err != nil {
    return
  }

  var id int
  var user_id int
  var character_id int

  row.Next()
  if err = row.Scan(&id, &user_id, &character_id); err != nil {
    return
  }

  collections = models.Collection {
    Id:          id,
    UserId:      user_id,
    CharacterId: character_id,
  }

  return
}
