package grand

import (
	"github.com/google/uuid"
)

type Grand struct {
	UID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string    `gorm:"type:varchar(100)"`
	Level int
}

func NewGrand(name string, level int) *Grand {
	return &Grand{
		UID:   uuid.New(),
		Name:  name,
		Level: level,
	}
}
