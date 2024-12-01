package character

import "github.com/google/uuid"

type Character struct {
	Id      *uuid.UUID `gorm:"primaryKey"`
	OwnerId *uuid.UUID `gorm:"index"`
}
