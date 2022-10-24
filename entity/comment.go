package entity

import (
	t "time"
)

type Comment struct {
	CommentID int `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt t.Time
  UpdatedAt t.Time
  Message string `validate:"required"`
	UserID int `gorm:"not null"`
	PhotoID int `gorm:"not null"`
}
