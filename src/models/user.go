package models

type User struct {
  Id int            `json:"id"`
  Name string       `json:"name"`
  Token string      `json:"token"`
  CreatedAt string  `json:"created_at"`
  UpdatedAt string  `json:"updated_at"`
}
