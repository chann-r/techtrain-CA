package database

import (
  "math/rand"
  "time"
  "log"
)

type GachaRepository struct {
	SqlHandler *SqlHandler
}

// timesの数だけランダムにcharacter_idを生成して返す
func (repo *GachaRepository) Choose(times int) (characterIds []int, err error){

	rows, err := repo.SqlHandler.Query("SELECT weight FROM probabilities ORDER BY id")
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
  
	for i := 1; i <= times; i++ { // 保存するキャラクターidの計算をtimesの数だけ実行
	  // 0から99までの乱数に1を足す(1から100までの乱数を生成)
	  percentage := rand.Intn(99) + 1
	  log.Print("The random percentage is ", percentage)
  
	  var character_id int // 保存するキャラクターid
  
	  // 保存するキャラクターidの計算
	  for j := 0; j < len(probabilities); j++ {
		if thresholds[j] <= percentage && percentage <= thresholds[j+1] {
		  character_id = j + 1
		}
	  }
	  log.Print("The character_id is ", character_id)
  
	  // 保存するキャラクターidを格納したスライス
	  characterIds = append(characterIds, character_id)
	}
  
	return
  }