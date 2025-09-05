package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/grand"
)

type User struct {
	UID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Token     string    `gorm:"type:varchar(100)"`
	GrandUID  uuid.UUID
	Grand     grand.Grand `gorm:"foreignKey:GrandUID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(token string, grandUID uuid.UUID) *User {
	return &User{
		UID:       uuid.New(),
		Token:     token,
		GrandUID:  grandUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
