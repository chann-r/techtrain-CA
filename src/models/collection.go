package models

type Collection struct {
  Id          int
  UserId      int
  CharacterId int
}

type Collections []Collection


type GachaDrawResponse struct {
  CharacterId int    `json:"characterID"`
  Name        string `json:"name"`
}

type GachaDrawResponses []GachaDrawResponse
