package database

import (
  "math/rand"
  "time"
)

type CollectionRepository struct {
  SqlHandler *SqlHandler
}

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
