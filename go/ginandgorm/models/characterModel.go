package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name     string
	RealName string
	Role     string
}

// equals
// type Character struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//  Name     string
//  RealName string
//  Role     string
// }
