package models

type Example struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
