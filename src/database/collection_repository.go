package database

import (
  "math/rand"
  "time"
  "techtrain-CA/models"
)

type CollectionRepository struct {
  SqlHandler *SqlHandler
}

// timesの数だけランダムにcharacter_idを生成してuser_idと一緒に保存する
func (repo *CollectionRepository) Store(user_id int, times int) (characterIds []int, err error) {
  // シードを与える（デフォルトだと同じ乱数ジェネレーターを使用してしまう）
  rand.Seed(time.Now().UnixNano())

  for i := 1; i <= times; i++ {
    // 0から2までの乱数に1を足す
    character_id := rand.Intn(2) + 1

    result, _ := repo.SqlHandler.Execute("INSERT INTO collections (user_id, character_id) VALUES (?, ?)", user_id, character_id)
    if err != nil {
      return
    }

    // 保存した行のidを取得
    identifier, _ := result.LastInsertId()
    if err != nil {
      return
    }

    // intに変換
    id := int(identifier)

    // スライスの要素に保存したcollectionのideを追加
    characterIds = append(characterIds, id)
  }

  return
}

// idを格納したスライスを引数に、それぞれのcollectionを検索して返す
func (repo *CollectionRepository) FindByIds(characterIds []int) (gachaDrawResponses models.GachaDrawResponses, err error) {
  for _, value := range characterIds {
    row, _ := repo.SqlHandler.Query("SELECT collections.character_id, characters.name from collections INNER JOIN characters ON collections.character_id = characters.id WHERE collections.id = ?", value)

    defer row.Close()

    if err != nil {
      return
    }

    var characterID int
    var name string

    row.Next()
    if err = row.Scan(&characterID, &name); err != nil {
      return
    }

    gachaDrawResponse := models.GachaDrawResponse {
      CharacterId: characterID,
      Name:        name,
    }

    gachaDrawResponses = append(gachaDrawResponses, gachaDrawResponse)
  }

  return
}

// collectionのidでcollectionを検索して返す
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
