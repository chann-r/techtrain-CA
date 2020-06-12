package models

type Collection struct {
  Id          int
  UserId      int
  CharacterId int
}

type Collections []Collection


type GachaDrawResponse struct {
  CharacterId string `json:"characterID"`
  Name        string `json:"name"`
}

type GachaDrawResponses []GachaDrawResponse


type UserCharacter struct {
  UserCharacterId string `json:"UserCharacterID"`
  CharacterId     string `json:"characterID"`
  Name            string `json:"name"`
}

type UserCharacters []UserCharacter
