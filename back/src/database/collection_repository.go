package database

import (
  "math/rand"
  "time"
  "strconv"
  "techtrain-CA/models"
  "log"
)

type CollectionRepository struct {
  SqlHandler *SqlHandler
}

// timesの数だけランダムにcharacter_idを生成して返す
func (repo *CollectionRepository) Choose(times int) (characterIds []int, err error){

  rows, err := repo.SqlHandler.Query("SELECT weight FROM probabilities")
  if err != nil {
    return
  }
  defer rows.Close()

  var probabilities []int // 確率を格納するスライス

  for rows.Next() { // 確率を格納するスライス
    var probability int

    if err = rows.Scan(&probability); err != nil {
      return
    }

    probabilities = append(probabilities, probability)
  }
  log.Print("The probability slice is ", probabilities)

  thresholds := [] int{0} //閾値を格納するスライス

  // 閾値を計算して格納
  for k := 0; k < len(probabilities); k++ {
    threshold := thresholds[k] + probabilities[k]

    thresholds = append(thresholds, threshold)
  }
  log.Print("The threshold slice is ", thresholds)

  // シードを与える（デフォルトだと同じ乱数ジェネレーターを使用してしまう）
  rand.Seed(time.Now().UnixNano())

  for i := 1; i <= times; i++ {
    // 0から99までの乱数に1を足す(1から100までの乱数を生成)
    percentage := rand.Intn(99) + 1
    log.Print("The random percentage is ", percentage)

    var character_id int // 保存するキャラクターid

    if thresholds[0] <= percentage && percentage <= thresholds[1] {
      character_id = 1
    } else if thresholds[1] <= percentage && percentage <= thresholds[2] {
      character_id = 2
    } else {
      character_id = 3
    }
    log.Print("The character_id is ", character_id)

    // 保存するキャラクターidを格納したスライス
    characterIds = append(characterIds, character_id)
  }

  return
}

// 選択されたキャラクターidとuser_idと一緒に保存する
func (repo *CollectionRepository) Store(user_id int, characterIds []int) (storedCharacterIds []int, err error) {

  for i := 1; i <= len(characterIds); i++ {
    result, _ := repo.SqlHandler.Execute("INSERT INTO collections (user_id, character_id) VALUES (?, ?)", user_id, characterIds[i-1])
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
    storedCharacterIds = append(storedCharacterIds, id)
  }

  return
}

// idを格納したスライスを引数に、それぞれのcollectionを検索して返す
func (repo *CollectionRepository) FindByIds(characterIds []int) (gachaDrawResponses models.GachaDrawResponses, err error) {
  for _, value := range characterIds {
    rows, _ := repo.SqlHandler.Query("SELECT collections.character_id, characters.name FROM collections INNER JOIN characters ON collections.character_id = characters.id WHERE collections.id = ?", value)

    defer rows.Close()

    if err != nil {
      return
    }

    var characterId int
    var name string

    rows.Next()
    if err = rows.Scan(&characterId, &name); err != nil {
      return
    }

    characterID := strconv.Itoa(characterId)

    gachaDrawResponse := models.GachaDrawResponse {
      CharacterId: characterID,
      Name:        name,
    }

    gachaDrawResponses = append(gachaDrawResponses, gachaDrawResponse)
  }

  return
}

// ユーザーidで所有キャラクターを検索して返す
func (repo *CollectionRepository) FindByUserId(user_id int) (userCharacters models.UserCharacters, err error) {
  rows, err := repo.SqlHandler.Query("SELECT collections.id, collections.character_id, characters.name FROM collections INNER JOIN characters ON collections.character_id = characters.id WHERE collections.user_id = ?", user_id)

  defer rows.Close()

  if err != nil {
    return
  }

  for rows.Next() {
    var userCharacterId int
    var characterId     int
    var name            string

    if err = rows.Scan(&userCharacterId, &characterId, &name); err != nil {
      return
    }

    userCharacterID := strconv.Itoa(userCharacterId)
    characterID     := strconv.Itoa(characterId)

    userCharacter := models.UserCharacter {
      UserCharacterId: userCharacterID,
      CharacterId:     characterID,
      Name:            name,
    }

    userCharacters = append(userCharacters, userCharacter)
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