package models

import "gorm.io/gorm"

type Account struct {
	*gorm.Model

	Username string `db:"username"`
	Avatar   string `db:"avatar"`
	FullName string `db:"full_name"`
}
