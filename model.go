package gorm

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// ModelSupportUnique add `deleted_flag` with Model
type ModelSupportUnique struct {
	Model
	DeletedFlag uint `gorm:"type:INT UNSIGNED DEFAULT 0 NOT NULL"`
}
