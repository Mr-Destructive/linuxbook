package models

type HelpBook struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name string `json:"name"`
  Description string `json:"description"`
  Parameters string `json:"parameters"`
}

type CreateHelpBookInput struct {
  Name string `json:"name" binding:"required"`
  Description string `json:"description" binding:"required"`
  Parameters string `json:"parameters" binding:"required"`
}

