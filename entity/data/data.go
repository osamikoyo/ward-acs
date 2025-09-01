package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/grand"
)

type Data struct {
	UID       uuid.UUID
	GrandUID  uuid.UUID
	Grand     grand.Grand `gorm:"foreignKey:GrandUID"`
	Payload   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Encrypted bool
}
